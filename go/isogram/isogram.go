//isogram package conatins functions that help us deal with isograms
package isogram

import (
	"strings"
	"unicode"
)

//IsIsogram takes a string as input and tells you if inputted string is an isogram
func IsIsogram(input string) bool {
	letters := make(map[rune]bool)
	for _, val := range strings.ToLower(input) {
		if !unicode.IsLetter(val) { //checking for non letter chracters, if found skip
			continue
		}
		if letters[val] { //if letter exists in map with true value it is a repeat
			return false
		} else {
			letters[val] = true //otherwise set letter to value true and add to map
		}
	}
	return true
}
