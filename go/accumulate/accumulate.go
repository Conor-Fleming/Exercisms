package accumulate

func Accumulate(given []string, operation func(string) string) []string {
	for i := range given {
		given[i] = operation(given[i])
	}
	return given
}
