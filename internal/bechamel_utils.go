package internal

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
