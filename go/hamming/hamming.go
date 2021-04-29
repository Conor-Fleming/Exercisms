//Hamming package provides functions to use for scientific and biological error checking
package hamming

import "fmt"

//Distance function takes two strings and checks that they are the same length before counting the differences between the strings
func Distance(a, b string) (int, error) {
	runesA, runesB := []rune(a), []rune(b) //changed strings to runes to account for utf-8 characters (thanks andre)
	if len(runesA) != len(runesB) {
		return 0, fmt.Errorf("Two values %q, and %q are not the same length", a, b)
	}
	var count int = 0
	for i := range runesA {
		if runesA[i] != runesB[i] {
			count += 1
		}
	}
	return count, nil
}
