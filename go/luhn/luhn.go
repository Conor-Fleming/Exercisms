package luhn

import (
	"strings"
	"unicode"
)

//Valid takes a SID number or credit card number and returns true or false based on its validity
func Valid(number string) bool {
	number = strings.ReplaceAll(number, " ", "")
	numbers := []rune(number)
	//checking that length is valid
	if len(numbers) < 2 {
		return false
	}
	//checking for non digit characters
	totalSum := 0
	for i := range numbers {
		//checking for non-digit chars
		if !unicode.IsDigit(numbers[i]) {
			return false
		}
		v := int(numbers[i] - '0')
		//determine which chars to double based on length of number
		if len(numbers)%2 == 0 {
			if i%2 == 0 {
				v = v * 2
				if v > 9 {
					v = v - 9
				}
			}
		} else {
			if i%2 != 0 {
				v = v * 2
				if v > 9 {
					v = v - 9
				}
			}
		}
		totalSum += v
	}
	return totalSum%10 == 0
}
