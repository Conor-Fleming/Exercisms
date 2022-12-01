package matrix

import (
	"strconv"
	"strings"
)

type Matrix [][]int

func New(s string) (*Matrix, error) {
	rows := strings.Split(s, " ")
	var matrix Matrix
	for i := 0; i < len(rows)/2; i++ {
		for y := 0; y < len(rows)/2; y++ {
			value, err := strconv.Atoi(rows[i])
			if err != nil {
				return nil, err
			}
			matrix[i][y] = value
		}
	}
	return &matrix, nil
}

// Cols and Rows must return the results without affecting the matrix.
func (m *Matrix) Cols() [][]int {
	return nil
}

func (m *Matrix) Rows() [][]int {
	return nil
}

func (m *Matrix) Set(row, col, val int) bool {
	return true
}
