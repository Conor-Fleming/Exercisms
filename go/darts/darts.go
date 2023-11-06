package darts

import "math"

func Score(x, y float64) int {
	point := math.Sqrt(x*x + y*y)

	switch {
	case point <= 1.0:
		return 10
	case point <= 5:
		return 5
	case point <= 10:
		return 1
	default:
		return 0
	}
}
