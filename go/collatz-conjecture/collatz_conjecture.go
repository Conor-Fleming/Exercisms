package collatzconjecture

import (
	"errors"
	"fmt"
)

var steps int = 0

func CollatzConjecture(n int) (int, error) {
	if n <= 0 {
		return 0, errors.New("zero or negative entry")
	}

	if n == 1 {
		fmt.Println("done")
		result := steps
		steps = 0
		return result, nil
	}

	steps += 1
	if n%2 == 0 {
		fmt.Println("even number", steps, ":", n)
		return CollatzConjecture(n / 2)
	} else {
		fmt.Println("odd", steps, ":", n)
		return CollatzConjecture((n * 3) + 1)
	}
}
