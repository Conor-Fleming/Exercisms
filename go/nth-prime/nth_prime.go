package prime

import (
	"errors"
	"fmt"
	"math"
)

// Nth returns the nth prime number. An error must be returned if the nth prime number can't be calculated ('n' is equal or less than zero)
func Nth(n int) (int, error) {
	if n <= 0 {
		return 0, errors.New("Nth prime cannot be computed, input was less than 1")
	}

	primes := []int{}
	i := 2
	for {
		if len(primes) == n {
			break
		}

		if IsPrime(float64(i)) {
			primes = append(primes, i)
		}
		i++
	}
	fmt.Println(primes)

	return primes[len(primes)-1], nil
}

func IsPrime(n float64) bool {
	if n <= 1 {
		return false
	}

	for i := 2; float64(i) < math.Sqrt(n); i++ {
		if int(n)%i == 0 {
			return false
		}
	}

	return true
}
