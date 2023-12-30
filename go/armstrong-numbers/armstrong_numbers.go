package armstrong

import (
	"math"
	"strconv"
)

func IsNumber(n int) bool {
	strN := strconv.Itoa(n)

	exponent := len(strN)

	result := 0.0
	for _, v := range strN {
		v, err := strconv.Atoi(string(v))
		if err != nil {
			panic(err)
		}

		result += math.Pow(float64(v), float64(exponent))
	}

	return int(result) == n
}
