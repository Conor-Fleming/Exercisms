package wordcount

import (
	"regexp"
	"strings"
)

type Frequency map[string]int

func WordCount(phrase string) Frequency {
	result := make(map[string]int)
	phrase = strings.ToLower(phrase)
	re := regexp.MustCompile(`\b[\w']+\b`)
	words := re.FindAllString(phrase, -1)
	for _, word := range words {
		result[word] += 1
	}
	return result
}
