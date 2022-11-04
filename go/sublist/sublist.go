package sublist

import "reflect"

// Relation type is defined in relations.go file.

func Sublist(l1, l2 []int) Relation {
	if reflect.DeepEqual(l1, l2) {
		return RelationEqual
	}
	if l1 == nil {
		return RelationSublist
	}
	if l2 == nil {
		return RelationSuperlist
	}

	var longArr []int
	var shortArr []int
	if len(l1) > len(l2) {
		longArr = l1
		shortArr = l2
	} else {
		longArr = l2
		shortArr = l1
	}

	i := 0
	y := 0
	for i < len(longArr) && y < len(shortArr) {
		if longArr[i] == shortArr[y] {
			i++
			y++
			if y == len(shortArr) {
				return RelationSublist
			}
		}
		i = i - y + 1
		y = 0
	}
	return RelationUnequal
}
