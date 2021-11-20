package main

import (
	"bufio"
	"fmt"
	"io"
	"neomantra/gotf/internal/gotf"
	"os"

	"github.com/spf13/pflag"
)

/////////////////////////////////////////////////////////////////////////////////////

var usageFormat string = `usage:  %s <options>

GOlang Time Formatter

Reads stdin, converts epoch times to human readable, outputs to stdout.

10-digits are interpreted as seconds, 13 as milliseconds,
16 as microseconds, and 19 as nanoseconds.

example:
  cat log.txt | gotf -g | head

`

/////////////////////////////////////////////////////////////////////////////////////
// Main Program

const DEFAULT_FORMAT = "15:04:05"
const DEFAULT_FORMAT_WITH_DATE = "2006010-15:04:05"

const BLOCK_BUFFER_SIZE = 4096

func main() {
	var outputFormat string
	var useDate bool
	var blockBuffering bool
	var globalMatch bool
	var showHelp bool

	pflag.StringVarP(&outputFormat, "format", "f", "", "golang Time.Format string (default: '15:04:05')")
	pflag.BoolVarP(&globalMatch, "global", "g", false, "global match")
	pflag.BoolVarP(&blockBuffering, "block", "b", false, "use block buffering (default: line buffering)")
	pflag.BoolVarP(&useDate, "date", "d", false, "default format with '2006010-15:04:05'")
	pflag.BoolVarP(&showHelp, "help", "h", false, "show help")
	pflag.Parse()

	if showHelp {
		fmt.Fprintf(os.Stdout, usageFormat, os.Args[0])
		pflag.PrintDefaults()
		os.Exit(0)
	}

	if len(outputFormat) == 0 {
		if useDate {
			outputFormat = DEFAULT_FORMAT_WITH_DATE
		} else {
			outputFormat = DEFAULT_FORMAT
		}
	}

	if blockBuffering {
		reader := bufio.NewReader(os.Stdin)
		buf := make([]byte, BLOCK_BUFFER_SIZE)
		for {
			n, err := reader.Read(buf)
			buf = buf[:n]
			if err != nil {
				if err == io.EOF {
					break
				}
				if err != io.ErrUnexpectedEOF {
					fmt.Fprintln(os.Stderr, err)
					break
				}
			}

			str, _ := gotf.ConvertTimes(string(buf), outputFormat, globalMatch)
			fmt.Println(str)
		}
	} else {
		// line buffering with bufio.Scanner
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			// convert time
			str, _ := gotf.ConvertTimes(scanner.Text(), outputFormat, globalMatch)
			fmt.Println(str)
		}

		if err := scanner.Err(); err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
		}
	}
}
