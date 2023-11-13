// Package bob lets our friend bob speak
package bob

import (
	"strings"
	"unicode"
)

// Hey allows bob to say hey
func Hey(remark string) string {
	//cleaning string
	remark = strings.TrimSpace(remark)

	switch {
	case remark == "":
		return "Fine. Be that way!"
	case isUpperWithLetters(remark) && strings.HasSuffix(remark, "?"):
		return "Calm down, I know what I'm doing!"
	case strings.HasSuffix(remark, "?"):
		return "Sure."
	case isUpperWithLetters(remark):
		return "Whoa, chill out!"
	default:
		return "Whatever."
	}
}

func isUpper(s string) bool {
	for _, v := range s {
		if !unicode.IsUpper(v) && unicode.IsLetter(v) {
			return false
		}
	}

	return true
}

func hasLetters(s string) bool {
	for _, v := range s {
		if unicode.IsLetter(v) {
			return true
		}
	}

	return false
}

func isUpperWithLetters(s string) bool {
	return isUpper(s) && hasLetters(s)
}
