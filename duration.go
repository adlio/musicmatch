package musicmeta

import (
	"strconv"
	"strings"
	"time"
)

// DurationFromString converts a song length in string
// format (e.g. 00:00) and returns a time.Duration.
//
func DurationFromString(durStr string) time.Duration {
	var d time.Duration
	increments := []time.Duration{time.Second, time.Minute, time.Hour}
	stringParts := strings.Split(durStr, ":")
	for i := len(stringParts) - 1; i >= 0; i-- {
		iPart, _ := strconv.Atoi(stringParts[i])
		j := len(stringParts) - 1 - i
		d += time.Duration(iPart) * increments[j]
	}
	return d
}

// DurationFromStringInSeconds converts a song length in string
// format (e.g. 00:00) and returns an integer representing the
// song's length in seconds.
//
func DurationFromStringInSeconds(durStr string) int {
	return int(DurationFromString(durStr).Seconds())
}
