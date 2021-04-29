package hamming

import "fmt"

func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, fmt.Errorf("Two values %q, and %q are not the same length", a, b)
	}
	var count int = 0
	for i := range a {
		if a[i] != b[i] {
			count += 1
		}
	}
	return count, nil
}
