package pangram

import (
	"strings"
)

func IsPangram(input string) bool {
	letters := make(map[string]int)

	input = strings.ToLower(input)
	input = strings.ReplaceAll(input, " ", "")

	for _, v := range input {
		key := string(v)
		if _, ok := letters[key]; !ok {
			letters[key] = 1
		}
	}
	if len(letters) == 26 {
		return true
	}
	return false
}
