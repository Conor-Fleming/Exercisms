package diffsquares

import "math"

func SquareOfSum(input int) int {
	result := 0
	for i := 0; i < input; i++ {
		result = i + result
	}
	return int(math.Pow(2, float64(result)))
}

func SumOfSquares(input int) int {

}

func Difference(SquareOfSum int, SumOfSquares int) int {

}
