package lsproduct

import (
	"errors"
	"strconv"
	"unicode"
)

func LargestSeriesProduct(digits string, span int) (int64, error) {
	// validating input values
	if span < 0 {
		return 0, errors.New("span must not be a negative number")
	}
	if len(digits) < span {
		return 0, errors.New("span must be smaller than string length")
	}
	if containsLetters(digits) {
		return 0, errors.New("digits input must only contain digits")
	}

	var max int

	for i := 0; i <= len(digits)-span; i++ {
		piece := digits[i : i+span]
		tempMax := findMax(piece)

		if tempMax > max {
			max = tempMax
		}
	}
	return int64(max), nil
}

func containsLetters(digits string) bool {
	for _, v := range digits {
		if unicode.IsLetter(v) {
			return true
		}
	}
	return false
}

func findMax(digits string) int {
	temp := 1
	for _, v := range digits {
		value, _ := strconv.Atoi(string(v))

		temp = temp * value
	}
	return temp
}
