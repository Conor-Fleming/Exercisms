package pangram

import (
	"strings"
	"unicode"
)

func IsPangram(input string) bool {
	letters := make(map[string]int)

	input = strings.ToLower(input)
	input = strings.ReplaceAll(input, " ", "")

	for _, v := range input {
		if _, ok := letters[string(v)]; !ok {
			if unicode.IsLetter(v) {
				letters[string(v)] = 1
			}
		}
	}

	if len(letters) == 26 {
		return true
	}
	return false
}
