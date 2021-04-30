//Scrabble provides way to take the math out of scrabble
package scrabble

import "strings"

//Score takes a word and tells you how much it is worth in a game of scrabble
func Score(word string) int {
	word = strings.ToUpper(word)
	score := 0
	for i := range word {
		switch word[i] {
		case 'A', 'E', 'I', 'O', 'U', 'L', 'N', 'R', 'S', 'T':
			score += 1
		case 'D', 'G':
			score += 2
		case 'B', 'C', 'M', 'P':
			score += 3
		case 'F', 'H', 'V', 'W', 'Y':
			score += 4
		case 'K':
			score += 5
		case 'J', 'X':
			score += 8
		case 'Z', 'Q':
			score += 10
		}
	}
	return score
}
