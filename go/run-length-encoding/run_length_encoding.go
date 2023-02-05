package encode

import (
	"strconv"
)

func RunLengthEncode(input string) string {
	var output string
	for i := 0; i < len(input); i++ {
		elem := input[i]
		count := 1
		for j := i + 1; j < len(input); j++ {
			if input[j] == elem {
				count++
				continue
			}
			i = j - 1
			break
		}
		if count == 1 {
			output += string(elem)
			continue
		}
		output += strconv.Itoa(count) + string(elem)
	}

	return output
}

func RunLengthDecode(input string) string {
	return ""
}
