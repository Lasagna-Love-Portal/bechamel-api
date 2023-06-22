package internal

import (
	"time"
)

// Project Ricotta: Bechamel API
//
// Internal utility functions used in Bechamel API

func StringIsInArray(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}

// NOTE: this is not an arbitrary formatting string, this is required format string
// to get time created in ISO 8601 simplified extended format as returned
// by JavaScript's toISOString() function.
func TimeAsISO8601String(theTime time.Time) string {
	return theTime.Format("2006-01-02T15:04:05.000Z0700")
}

func CurrentTimeAsISO8601String() string {
	return TimeAsISO8601String(time.Now().UTC())
}
