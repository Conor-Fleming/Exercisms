package grains

import (
	"errors"
)

func Square(input int) (uint64, error) {
	if input < 1 || input > 64 {
		return 0, errors.New("input must be greater than 0 and less than 65")
	}
	return (1 << uint64(input-1)), nil
}

func Total() uint64 {
	return 1<<64 - 1
}
