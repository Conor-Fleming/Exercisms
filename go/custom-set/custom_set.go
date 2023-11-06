package stringset

import "fmt"

// Implement Set as a collection of unique string values.
//
// For Set.String, use '{' and '}', output elements as double-quoted strings
// safely escaped with Go syntax, and use a comma and a single space between
// elements. For example, a set with 2 elements, "a" and "b", should be formatted as {"a", "b"}.
// Format the empty set as {}.

// Define the Set type here.
type Set struct {
	val []string
}

func New() Set {
	return Set{}
}

func NewFromSlice(l []string) Set {
	l = removeDuplicates(l)

	return Set{
		val: l,
	}
}

func (s Set) String() string {
	var output string
	output += "{"
	for i, v := range s.val {
		if i == len(s.val)-1 {
			output += fmt.Sprintf(`"%v"`, v)
			continue
		}
		output += fmt.Sprintf(`"%v", `, v)
	}
	output += "}"

	return output
}

func (s Set) IsEmpty() bool {
	return len(s.val) == 0
}

func (s Set) Has(elem string) bool {
	for _, v := range s.val {
		if v == elem {
			return true
		}
	}

	return false
}

func (s *Set) Add(elem string) {
	s.val = append(s.val, elem)
}

func Subset(s1, s2 Set) bool {
	seen := make(map[string]bool)
	for _, v := range s2.val {
		seen[v] = true
	}

	for _, item := range s1.val {
		if !seen[item] {
			return false
		}
	}

	return true
}

func Disjoint(s1, s2 Set) bool {
	for _, v := range s2.val {
		if s1.Has(v) {
			return false
		}
	}

	return true
}

func Equal(s1, s2 Set) bool {
	for _, v := range s2.val {
		if !s1.Has(v) {
			return false
		}
	}

	for _, v := range s1.val {
		if !s2.Has(v) {
			return false
		}
	}

	return true
}

func Intersection(s1, s2 Set) Set {
	var result = Set{}

	for _, v := range s1.val {
		if s2.Has(v) {
			result.val = append(result.val, v)
		}
	}

	return result
}

func Difference(s1, s2 Set) Set {
	result := Set{}
	for _, v := range s1.val {
		if !s2.Has(v) {
			result.val = append(result.val, v)
		}
	}

	return result
}

func Union(s1, s2 Set) Set {
	result := s1
	for _, v := range s2.val {
		if !result.Has(v) {
			result.val = append(result.val, v)
		}
	}

	return result
}

func removeDuplicates(slice []string) []string {
	seen := make(map[string]bool)
	result := []string{}

	for _, item := range slice {
		if !seen[item] {
			seen[item] = true
			result = append(result, item)
		}
	}

	return result
}
