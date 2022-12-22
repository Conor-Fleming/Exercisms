package matrix

import (
	"strconv"
	"strings"
)

type Matrix [][]int

func New(s string) (*Matrix, error) {
	//need to do some error checking for:
	// int overflows (long integers), uneven rows, empty rows, non ints and non numerics
	//basic string manipulation to enable easy parsing
	s = strings.ReplaceAll(s, "\n", " . ")
	rowCount := strings.Count(s, ".")
	lines := strings.Split(s, " ")

	//create matrix so store values
	mx := make(Matrix, rowCount+1)

	i := 0
	//loop range of string and add to matrix row
	for _, val := range lines {
		// the current value is the newline marker (".") then increment matrix row with i++
		if val == "." {
			i++
			continue
		}
		val, _ := strconv.Atoi(string(val))
		mx[i] = append(mx[i], val)
	}

	return &mx, nil
}

// Cols and Rows must return the results without affecting the matrix.
func (m *Matrix) Cols() [][]int {
	return nil
}

func (m *Matrix) Rows() [][]int {
	return nil
}

func (m *Matrix) Set(row, col, val int) bool {
	return false
}
