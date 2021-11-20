package gotf

import (
	"regexp"
	"strconv"
	"time"
)

// 10-digits epoch
// 13-digits epoch+3 milliseconds
// 16-digits epoch+6 microseconds
// 19-digits epoch+9 nanoseconds
var timevalRegex = regexp.MustCompile(`([0-9]{19}|[0-9]{16}[0-9]{13}|[0-9]{10})`)

// Converts an "epoch" string to a Time.
// 10-digits are interpreted as seconds, 13 as milliseconds,
// 16 as microseconds, and 19 as nanoseconds
// Returns (time.Time{}, error) on error.
func EpochToTime(str string) (time.Time, error) {
	num, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return time.Time{}, err
	}
	switch len(str) {
	case 13:
		sec, msec := num/1000, num%1000
		return time.Unix(sec, msec*1000000), nil
	case 16:
		sec, usec := num/1000000, num%1000000
		return time.Unix(sec, usec*1000), nil
	case 19:
		sec, nsec := num/1000000000, num%1000000000
		return time.Unix(sec, nsec), nil
	default:
		return time.Unix(num, 0), nil
	}
}

// Converts all time-like strings found in `str` to the supplied Time.Format string
// Returns the converted string and true if the original string was modified.
func ConvertTimes(str string, outFormat string, globalMatch bool) (string, bool) {
	everConverted := false
	res := timevalRegex.ReplaceAllStringFunc(str, func(part string) string {
		if !globalMatch && everConverted {
			return part
		}
		if tv, err := EpochToTime(part); err != nil {
			return part
		} else {
			everConverted = true
			return tv.Format(outFormat)
		}
	})
	return res, everConverted
}
