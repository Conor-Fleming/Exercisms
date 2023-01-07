package strain

type Ints []int
type Lists [][]int
type Strings []string

func (ints Ints) Keep(filter func(int) bool) Ints {
	var result Ints
	for _, v := range ints {
		if filter(v) {
			result = append(result, v)
		}
	}
	return result
}

func (ints Ints) Discard(filter func(int) bool) Ints {
	var result Ints
	for _, v := range ints {
		if !filter(v) {
			result = append(result, v)
		}
	}
	return result
}

func (lists Lists) Keep(filter func([]int) bool) Lists {
	var result Lists
	for _, v := range lists {
		if filter(v) {
			result = append(result, v)
		}
	}
	return result
}

func (strings Strings) Keep(filter func(string) bool) Strings {
	var result Strings
	for _, v := range strings {
		if filter(v) {
			result = append(result, v)
		}
	}
	return result
}
