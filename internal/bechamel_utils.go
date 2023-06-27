package internal

import (
	"reflect"
	"strings"
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

/*
	    Verify that a provided value is of a suitable type to be assigned to a field (assignee)
		in a PATCH operation

		We consider float64 to be assignable to int and vice-versa, as the Bechamel API interface
		currently only has integer numeric values. These come in from the JSON as float64 types.
*/
func ValuesAreUpdateCompatible(assignee reflect.Value, value reflect.Value) bool {
	if assignee.Type() == value.Type() {
		return true
	}
	if strings.HasPrefix(assignee.Type().String(), "int") &&
		strings.HasPrefix(value.Type().String(), "float") {
		return true
	}
	if strings.HasPrefix(assignee.Type().String(), "float") &&
		strings.HasPrefix(value.Type().String(), "int") {
		return true
	}
	// If interface to array, verify types in there match those of the assignee
	if strings.HasPrefix(assignee.Type().String(), "[]") &&
		strings.HasPrefix(value.Type().String(), "[]interface") {
		for i := 0; i < value.Len(); i++ {
			var innerValue = value.Index(i)
			if assignee.Type().String()[2:] != innerValue.Elem().Type().String() {
				return false
			}
		}
		return true
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
