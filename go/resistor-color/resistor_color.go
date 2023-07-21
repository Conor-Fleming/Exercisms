package resistorcolor

type Color struct {
	color string
	code  int
}

var colors = []Color{
	{"black", 0},
	{"brown", 1},
	{"red", 2},
	{"orange", 3},
	{"yellow", 4},
	{"green", 5},
	{"blue", 6},
	{"violet", 7},
	{"grey", 8},
	{"white", 9},
}

// Colors should return the list of all colors.
func Colors() []string {
	result := []string{}
	for _, val := range colors {
		result = append(result, val.color)
	}
	return result
}

// ColorCode returns the resistance value of the given color.
func ColorCode(color string) int {
	for _, v := range colors {
		if color == v.color {
			return v.code
		}
	}
	return 0
}
