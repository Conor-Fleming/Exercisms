package reverse

func Reverse(s string) string {
	runeString := []rune(s)
	var result string
	for i := len(runeString) - 1; i >= 0; i-- {
		result += string(runeString[i])
	}

	return result
}
