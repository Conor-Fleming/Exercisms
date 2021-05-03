package diffsquares

func SquareOfSum(input int) int {
	/*result := 0
	for i := 0; i <= input; i++ {
		result = i + result
	}
	return result * result*/

	//the above commented code is equivalent to the following formula
	return input * input * (1 + input) * (1 + input) / 4
}

func SumOfSquares(input int) int {
	/*result := 0
	for i := 0; i <= input; i++ {
		sqr := i * i
		result += sqr
	}
	return result*/

	//the above commented code is equivalent to the following formula
	return input * (input + 1) * (2*input + 1) / 6
}

func Difference(input int) int {
	return SquareOfSum(input) - SumOfSquares(input)
}
