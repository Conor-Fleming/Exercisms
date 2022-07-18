package wordcount

import (
	"strings"
)

type Frequency map[string]int

func WordCount(phrase string) Frequency {
	result := make(map[string]int)
	noPunc := strings.NewReplacer(",", ".", ";", "/")
	phrase = noPunc.Replace(phrase)
	words := strings.Fields(phrase)
	for _, word := range words {
		if _, ok := result[word]; ok {
			result[word] = 1
		} else {
			result[word]++
		}
	}
	return result
}
