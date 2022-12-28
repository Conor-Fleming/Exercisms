package matrix

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Matrix [][]int

func New(s string) (*Matrix, error) {
	//basic error checking for invalid entrys
	if s[0] == '\n' || s[len(s)-1] == '\n' {
		return nil, errors.New("Invalid Entry")
	}

	if !strings.Contains(s, " ") && !strings.Contains(s, "\n") && len(s) > 1 {
		return nil, errors.New("Invalid Entry")
	}

	s = strings.ReplaceAll(s, "\n", " ! ")
	rowCount := strings.Count(s, "!")
	lines := strings.Split(s, " ")
	for i, v := range lines {
		if v == "!" {
			if lines[i+1] == "!" || lines[i-1] == "!" {
				return nil, errors.New("Invalid Entry")
			}
		}
	}

	//create matrix so store values
	mx := make(Matrix, rowCount+1)

	i := 0
	//loop range of string and add to matrix row
	for _, val := range lines {
		// the current value is the newline marker (".") then increment matrix row with i++
		if val == "!" {
			i++
			continue
		}
		val, _ := strconv.Atoi(string(val))
		mx[i] = append(mx[i], val)
	}

	length := len(mx[0])
	for _, v := range mx {
		if length != len(v) {
			return nil, errors.New("Uneven Rows")
		}
	}

	return &mx, nil
}

// Cols and Rows must return the result without affecting the matrix.
func (m *Matrix) Cols() [][]int {
	lenCheck := len(*m)
	fmt.Println(m)
	colsRes := make([][]int, lenCheck)
	fmt.Println(colsRes)
	for i, v := range *m {
		for y := range v {
			colsRes[i] = append(colsRes[i], (*m)[y][i])
		}
	}

	return colsRes
}

func (m *Matrix) Rows() [][]int {
	lenCheck := len(*m)
	fmt.Println(m)
	rowRes := make([][]int, lenCheck)
	fmt.Println(rowRes)
	for i, v := range *m {
		for y := range v {
			rowRes[i] = append(rowRes[i], (*m)[i][y])
		}
	}

	return rowRes
}

func (m *Matrix) Set(row, col, val int) bool {
	(*m)[row][col] = val
	return true
}
