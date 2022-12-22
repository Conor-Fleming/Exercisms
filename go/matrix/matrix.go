package matrix

import (
	"strconv"
	"strings"
)

type Matrix [][]int

func New(s string) (*Matrix, error) {
	s = strings.ReplaceAll(s, "\n", " . ")
	rowCount := strings.Count(s, ".")
	lines := strings.Split(s, " ")
	mx := make(Matrix, rowCount+1)

	i := 0
	for _, val := range lines {
		if val == "." {
			i++
			continue
		}
		val, _ := strconv.Atoi(string(val))
		mx[i] = append(mx[i], val)
	}

	//fmt.Println(mx)

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
