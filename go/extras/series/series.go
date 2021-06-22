package series

//All function returns a list of all substrings of s with length n
func All(n int, s string) []string {
	var result []string

	return result
}

//UnsafeFirst func returns the first substring of s with length n
func UnsafeFirst(n int, s string) string {
	var result string
	for i := 0; i < n; i++ {
		result = result + string(s[i])
	}
	return result
}
