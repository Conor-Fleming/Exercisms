package anagram

import (
	"sort"
	"strings"
)

func Detect(subject string, candidates []string) []string {
	anagrams := make([]string, 0)
	subject = strings.ToLower(subject)
	subjectSorted := sortString(subject)
	for _, v := range candidates {
		val := strings.ToLower(v)
		if subject == val {
			continue
		}
		if subjectSorted == sortString(val) {
			anagrams = append(anagrams, v)
		}
	}

	return anagrams
}

func sortString(s string) string {
	runes := []rune(s)
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] > runes[j]
	})

	return string(runes)
}
