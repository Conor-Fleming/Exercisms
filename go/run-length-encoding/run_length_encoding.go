package encode

import (
	"strconv"
)

func RunLengthEncode(input string) string {
	var output string
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
	}

	return output
}

func RunLengthDecode(input string) string {
	return ""
}
