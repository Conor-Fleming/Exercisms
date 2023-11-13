// Package bob lets our friend bob speak
package bob

import (
	"strings"
	"unicode"
)

// Hey allows bob to say hey
func Hey(remark string) string {
	switch {
	case strings.TrimSpace(remark) == "":
		return "Fine. Be that way!"
	case strings.HasSuffix(remark, "?"):
		return "Sure."
	case isUpper(remark) && strings.HasSuffix(remark, "?"):
		return "Calm down, I know what I'm doing!"
	case isUpper(remark):
		return "Whoa, chill out!"
	default:
		return "Whatever."
	}
}

func isUpper(s string) bool {
	for _, v := range s {
		if !unicode.IsUpper(v) {
			return false
		}
	}
	return true
}
