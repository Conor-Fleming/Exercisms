package etl

import "strings"

func Transform(in map[int][]string) map[string]int {
	newMap := make(map[string]int)
	for key, val := range in {
		for _, v := range val {
			v = strings.ToLower(v)
			newMap[v] = key
		}
	}
	return newMap
}
