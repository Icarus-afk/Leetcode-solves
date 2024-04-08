package main

import (
	"fmt"
)

func isValidSudoku(board [][]byte) bool {
	rows := [9][9]bool{}
	cols := [9][9]bool{}
	boxes := [9][9]bool{}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				num := board[i][j] - '1'
				boxIndex := (i/3)*3 + j/3
				if rows[i][num] || cols[j][num] || boxes[boxIndex][num] {
					return false
				}
				rows[i][num] = true
				cols[j][num] = true
				boxes[boxIndex][num] = true
			}
		}
	}
	return true
}

func main() {
	board := make([][]byte, 9)
	for i := range board {
		board[i] = make([]byte, 9)
		for j := range board[i] {
			fmt.Scan(&board[i][j])
		}
	}

	fmt.Println(isValidSudoku(board))
}
