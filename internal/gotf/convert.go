package gotf

import (
	"regexp"
	"strconv"
	"time"
)

var timevalRegex = regexp.MustCompile(`([0-9]{10})`)

// Converts an "epoch" string to a time.
// Returns nil, error on error.
func EpochToTime(s string) (time.Time, error) {
	sec, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return time.Time{}, err
	}
	return time.Unix(sec, 0), nil
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
