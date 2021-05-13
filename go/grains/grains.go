//Grains package provides functions for grains and chessboard info
package grains

import (
	"errors"
)

//Square function takes an integer and returns the number of grains on the corresponding square
func Square(input int) (uint64, error) {
	if input < 1 || input > 64 {
		return 0, errors.New("input must be greater than 0 and less than 65")
	}
	return (1 << (input - 1)), nil
}

//Total function gives you the total number of grains on a chess board
func Total() uint64 {
	return 1<<64 - 1
}
