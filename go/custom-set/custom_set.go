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
	return Set{
		val: l,
	}
}

func (s Set) String() string {
	var output string
	output += "{"
	for i, v := range s.val {
		if i >= len(s.val)-2 {
			output += fmt.Sprintf(`"%v"`, v)
			continue
		}
		output += fmt.Sprintf(`"%v, "`, v)
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
	panic("Please implement the Subset function")
}

func Disjoint(s1, s2 Set) bool {
	panic("Please implement the Disjoint function")
}

func Equal(s1, s2 Set) bool {
	panic("Please implement the Equal function")
}

func Intersection(s1, s2 Set) Set {
	panic("Please implement the Intersection function")
}

func Difference(s1, s2 Set) Set {
	panic("Please implement the Difference function")
}

func Union(s1, s2 Set) Set {
	panic("Please implement the Union function")
}
