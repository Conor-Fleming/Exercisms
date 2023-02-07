package encode

import (
	"strconv"
)

func RunLengthEncode(input string) string {
	var output string
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
	}

	return output
}

func RunLengthDecode(input string) string {
	return ""
}
