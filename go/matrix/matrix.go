package matrix

import (
	"fmt"
	"strconv"
	"strings"
)

type Matrix [][]int

func New(s string) (*Matrix, error) {
	rows := strings.Split(s, " ")
	rows = append(rows, "\n")
	fmt.Println(rows)

	mx := make(Matrix, 0)

	temp := make([]int, 0)
	for _, v := range rows {
		if v != "\n" {
			value, err := strconv.Atoi(v)
			if err != nil {
				return nil, err
			}
			temp = append(temp, value)
			continue
		}
		v = ""
		mx = append(mx, temp)
		temp = nil
	}
	fmt.Println(mx)

	return nil, nil
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
