package diffsquares

func SquareOfSum(input int) int {
	result := 0
	for i := 0; i <= input; i++ {
		result = i + result
	}
	return result * result
}

func SumOfSquares(input int) int {
	result := 0
	for i := 0; i <= input; i++ {
		sqr := i * i
		result += sqr
	}
	return result
}

func Difference(input int) int {
	return SquareOfSum(input) - SumOfSquares(input)
}
