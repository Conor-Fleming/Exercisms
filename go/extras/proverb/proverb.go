// Package proverb should have a package comment that summarizes what it's about.
package proverb

import "fmt"

// Proverb should have a comment documenting it.
func Proverb(rhyme []string) []string {
	if len(rhyme) == 0 { //check empty input slice
		return rhyme
	}
	output := []string{} //create slice to store output
	if len(rhyme) > 1 {
		for i := 0; i < len(rhyme)-1; i++ {
			output = append(output, fmt.Sprintf("For want of a %s the %s was lost.", rhyme[i], rhyme[i+1]))
		}
	}
	output = append(output, fmt.Sprintf("And all for the want of a %s.", rhyme[0]))
	return output
}
