package romannumerals

func ToRomanNumeral(input int) (string, error) {
	result := ""
	if input > 1000 {
		q, r := input/1000, input%1000
		for i := 0; i < q; i++ {
			result += "M"
		}
	}
}
