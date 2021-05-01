//isogram package conatins functions that help us deal with isograms
package isogram

import (
	"strings"
	"unicode"
)

//IsIsogram takes a string as input and tells you if inputted string is an isogram
func IsIsogram(input string) bool {
	input = strings.Replace(input, "-", "", -1) //removing dashes
	input = strings.Replace(input, " ", "", -1) //removing spaces
	input = strings.ToLower(input)              //setting string to lower case
	letters := make(map[rune]bool)
	//fmt.Println(letters)
	for _, val := range input {
		if !unicode.IsLetter(val) { //checking for non letter chracters, if found skip
			continue
		}
		if letters[val] { //if letter exists in map with true value it is a repeat
			return false
		} else {
			letters[val] = true //otherwise set letter to value true and add to map
		}
	}
	//fmt.Println(letters)
	return true
}
