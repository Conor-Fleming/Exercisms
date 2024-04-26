package encode

import (
	"strconv"
)

func RunLengthEncode(input string) string {
	var output string
<<<<<<< HEAD
	boundCheck := len(input) - 1
	flag := false

	for i := 0; i < len(input); i++ {
		count := 1
		check := input[i]

		for j := i + 1; j < len(input); j++ {
			if j >= boundCheck {
				flag = true
			}
			if input[j] != check {
				i = j - 1
				break
			}
			count++
		}
		if count == 1 {
			output += string(check)
		} else {
			output += strconv.Itoa(count) + string(check)
		}

		if flag {
			break
		}
=======
	count := 1
	for i, v := range input {
		if i < len(input)-1 && byte(v) == input[i+1] {
			count += 1
			continue
		}
		if count > 1 {
			output += strconv.Itoa(count)
		}
		output += string(v)

		//reset counter upon finding a non match
		count = 1
>>>>>>> df2a20f0eb81f848c3148ec35c93c31da0b57d52
	}

	return output
}

func RunLengthDecode(input string) string {
	return ""
}
