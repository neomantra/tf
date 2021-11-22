package main

import (
	"bufio"
	"fmt"
	"io"
	"neomantra/tf/internal/tf"
	"os"

	"github.com/spf13/pflag"
)

/////////////////////////////////////////////////////////////////////////////////////

var usageFormat string = `usage:  %s <options> [file1 [file2 [...]]]

GOlang Time Formatter

Reads text files, converting epoch times to human readable, outputting to stdout.

10-digits are interpreted as seconds, 13 as milliseconds,
16 as microseconds, and 19 as nanoseconds.

If no filenames or only '-' is passed, stdin is processed.

example:
echo 1637421447 | tf

tf -g log.txt | head

`

////////////////////////////////////////////////////////////////////////////////////
// Globals

var g_outputFormat string = ""
var g_blockBuffering bool = false
var g_globalMatch bool = false

////////////////////////////////////////////////////////////////////////////////////

func processReader(reader io.Reader) error {
	if g_blockBuffering {
		// block buffering with bufio Reader
		reader := bufio.NewReader(reader)
		buf := make([]byte, BLOCK_BUFFER_SIZE)
		for {
			n, err := reader.Read(buf)
			buf = buf[:n]
			if err != nil {
				if err == io.EOF {
					return nil
				}
				if err != io.ErrUnexpectedEOF {
					return err
				}
			}

			str, _ := tf.ConvertTimes(string(buf), g_outputFormat, g_globalMatch)
			fmt.Println(str)
		}
	} else {
		// line buffering with bufio.Scanner
		scanner := bufio.NewScanner(reader)
		for scanner.Scan() {
			// convert time
			str, _ := tf.ConvertTimes(scanner.Text(), g_outputFormat, g_globalMatch)
			fmt.Println(str)
		}

		if err := scanner.Err(); err != nil {
			return err
		}
	}
	return nil
}

/////////////////////////////////////////////////////////////////////////////////////
// Main Program

const DEFAULT_FORMAT = "15:04:05"
const DEFAULT_FORMAT_WITH_DATE = "2006-01-02 15:04:05"

const BLOCK_BUFFER_SIZE = 4096

func main() {
	var useDate bool
	var showHelp bool

	pflag.StringVarP(&g_outputFormat, "format", "f", "", "golang Time.Format string (default: '15:04:05')")
	pflag.BoolVarP(&g_globalMatch, "global", "g", false, "global match")
	pflag.BoolVarP(&g_blockBuffering, "block", "b", false, "use block buffering (default: line buffering)")
	pflag.BoolVarP(&useDate, "date", "d", false, "default format with '2006-01-02 15:04:05'")
	pflag.BoolVarP(&showHelp, "help", "h", false, "show help")
	pflag.Parse()

	if showHelp {
		fmt.Fprintf(os.Stdout, usageFormat, os.Args[0])
		pflag.PrintDefaults()
		os.Exit(0)
	}

	if len(g_outputFormat) == 0 {
		if useDate {
			g_outputFormat = DEFAULT_FORMAT_WITH_DATE
		} else {
			g_outputFormat = DEFAULT_FORMAT
		}
	}

	// figure out files
	filenames := pflag.Args()
	if len(filenames) == 0 || (len(filenames) == 1 && filenames[0] == "-") {
		if err := processReader(os.Stdin); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		return
	}
	for _, filename := range filenames {
		f, err := os.Open(filename)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		defer f.Close()
		if err := processReader(f); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}
}
