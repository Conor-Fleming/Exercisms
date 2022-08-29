package pangram

import "strings"

func IsPangram(input string) bool {
	input = strings.ToLower(input)
	letters := make(map[string]int)

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
