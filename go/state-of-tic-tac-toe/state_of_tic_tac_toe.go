package main

import "fmt"

type State string

const (
	Win     State = "win"
	Ongoing State = "ongoing"
	Draw    State = "draw"
)

func main() {
	test := []string{"XOO", "X  ", "X  "}
	StateOfTicTacToe(test)
}

func StateOfTicTacToe(input []string) (State, error) {
	board := [][]rune{{' ', ' ', ' '}, {' ', ' ', ' '}, {' ', ' ', ' '}}

	for i, val := range input {
		for y, valy := range val {
			board[y][i] = valy
		}
	}

	return "", nil
}

func printBoard(board [][]string) {
	fmt.Printf("%v %v %v\n", string(board[0][0]), string(board[0][1]), string(board[0][2]))
	fmt.Printf("%v %v %v\n", string(board[1][0]), string(board[1][1]), string(board[1][2]))
	fmt.Printf("%v %v %v\n", string(board[2][0]), string(board[2][1]), string(board[2][2]))
}
func createMatrix(input []string) [][]rune {
	board := [][]rune{{' ', ' ', ' '}, {' ', ' ', ' '}, {' ', ' ', ' '}}

	for i, val := range input {
		for y, valy := range val {
			board[y][i] = valy
		}
	}
	return board
}
