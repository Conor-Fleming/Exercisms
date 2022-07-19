package wordcount

import (
	"strings"
)

type Frequency map[string]int

func WordCount(phrase string) Frequency {
	phrase = strings.ToLower(phrase)
	phrase = strings.TrimRight(phrase, "'")
	result := make(map[string]int)
	noPunc := strings.NewReplacer("' ", " ", " '", " ", ",", " ", ".", " ", ";", " ", "/", " ", ":", "", "!", "", "&", "", "@", "", "$", "", "%", "", "^", "", "' ", "")
	phrase = noPunc.Replace(phrase)
	words := strings.Fields(phrase)
	for _, word := range words {
		if result[word] == 0 {
			result[word] = 1
		} else {
			result[word]++
		}
	}
	return result
}
