package main

import "github.com/Dominic-Santos/gotest/utils"

func MoveNumberToRowCol(move int, board [][]string) (int, int) {
	current := 0
	for row := 0; row < len(board); row++ {
		for col := 0; col < len(board); col++ {
			if board[row][col] == " " {
				current++
				if current == move {
					return row, col
				}
			}
		}
	}
    return 0, 0
}

func FindMoveRandom(board [][]string) (int, int) {
	available := NumberAvailableMoves(board)
	if available != 0 {
		return MoveNumberToRowCol(utils.RandomNumber(1, available), board)
	}
	return 0, 0
}

func AiTurn (mode int, board [][]string) {
	var row, col int
	switch mode {
	default:
		row, col = FindMoveRandom(board)
	}
	
    PlayMove(row, col, false, board)
}
