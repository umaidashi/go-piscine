package piscine

import "ft"

func EightQueens() {
	var board [8]int
	solve(board, 0)
}

func solve(board [8]int, row int) {
	if row == 8 {
		printSolution(board)
		return
	}
	for i := 0; i < 8; i++ {
		if isSafe(board, row, i) {
			board[row] = i
			solve(board, row+1)
		}
	}
}

func printSolution(board [8]int) {
	for i := 0; i < 8; i++ {
		ft.PrintRune(rune(board[i]) + '1')
	}
	ft.PrintRune('\n')
}

func isSafe(board [8]int, row int, col int) bool {
	for i := 0; i < row; i++ {
		if board[i] == col || board[i]-col == i-row || board[i]-col == row-i {
			return false
		}
	}
	return true
}
