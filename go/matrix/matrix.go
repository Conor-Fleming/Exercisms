package matrix

type Matrix [][]int

func New(s string) (*Matrix, error) {
	var matrix Matrix
	for i, val := range s {
		y := 0
		if val == '\n' {
			i++
			continue
		}
		matrix[y][i] = int(val)
	}
	return &matrix, nil
}

// Cols and Rows must return the results without affecting the matrix.
func (m *Matrix) Cols() [][]int {
	panic("Please implement the Cols function")
}

func (m *Matrix) Rows() [][]int {
	panic("Please implement the Rows function")
}

func (m *Matrix) Set(row, col, val int) bool {
	panic("Please implement the Set function")
}
