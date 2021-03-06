package main

import (
	"github.com/Dominic-Santos/gotest/utils"
)

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

func FindRandomMove(board [][]string) (int, int) {
	available := NumberAvailableMoves(board)
	if available != 0 {
		return MoveNumberToRowCol(utils.RandomNumber(1, available), board)
	}
	return 0, 0
}

func FindWinningPos(arr []string, lookfor string) int {
	found := -1
	for i := 0; i < len(arr); i++ {
		if arr[i] != lookfor && arr[i] != " "  {
			return -1
		}
		if arr[i] == " " {
			if found == -1 {
				found = i
			} else {
				return -1
			}
		}
	}
	return found
}

func FindWinningMove(board [][]string, lookfor string) (bool, int, int){
	// check rows for winning move
	for row := 0; row < len(board); row++ {
		col := FindWinningPos(board[row], lookfor)
		if col != -1 {
			return true, row, col
		}
	}
	// check cols for winning move
	for col := 0; col < len(board); col++ {
		var cur []string
		for row := 0; row < len(board); row++ {
			cur = append(cur, board[row][col])
		}
		row := FindWinningPos(cur, lookfor)
		if row != -1 {
			return true, row, col
		}
	}
	// check diags for winning move
	var topleft []string
	var topright []string
	for i := 0; i < len(board); i++ {
		topleft = append(topleft, board[i][i])
		topright = append(topright, board[i][len(board) - 1 - i])
	}
	res := FindWinningPos(topleft, lookfor)
	if res != -1 {
		return true, res, res
	}
	res = FindWinningPos(topright, lookfor)
	if res != -1 {
		return true, res, len(board) - 1 - res
	}
	return false, 0, 0
}

func FindUrgentMove(board [][]string) (bool, int, int) {
	var status bool
	var row, col int
	// try to win
	status, row, col = FindWinningMove(board, "X")
	if status {
		return true, row, col
	}
	// try not to lose
	status, row, col = FindWinningMove(board, "O")
	if status {
		return true, row, col
	}
	return false, 0, 0
}

func FindGoodMove(board [][]string) (int, int) {
	urgent, row, col := FindUrgentMove(board)
	if urgent {
		return row, col
	}
	return FindRandomMove(board)
}

func CalcStrength(row, col int, board [][]string) (int, int) {
	var influence, length, tmplen int
	// check row open
	state := true
	for i := 0; i < len(board); i++ {
		switch board[row][i] {
		case "O":
			state = false
			break
		case "X":
			tmplen++
		}
	}
	if state {
		influence++
		length = tmplen
	}
	// check column open
	state = true
	tmplen = 0
	for i := 0; i < len(board); i++ {
		switch board[i][col] {
		case "O":
			state = false
			break
		case "X":
			tmplen++
		}
	}
	if state {
		influence++
		length = utils.MaxInt(length, tmplen)
	}
	// check topleft diagonal
	if row == col {
		state = true
		tmplen = 0
		for i := 0; i < len(board); i++ {
			switch board[i][i] {
			case "O":
				state = false
				break
			case "X":
				tmplen++
			}
		}
		if state {
			influence++
			length = utils.MaxInt(length, tmplen)
		}
	}

	// check topright diagonal
	if row == len(board) - 1 - col {
		state = true
		tmplen = 0
		for i := 0; i < len(board); i++ {
			switch board[i][len(board) - 1 - i] {
			case "O":
				state = false
				break
			case "X":
				tmplen++
			}
		}
		if state {
			influence++
			length = utils.MaxInt(length, tmplen)
		}
	}

	return influence, length
}

func FindSmartestMove(board [][]string) (int, int) {
	var bestRow, bestCol, bestInfluence, bestLength, equalCount int
	for i := 0; i < len(board) * len(board); i++ {
		if (board[i / len(board)][i % len(board)] == " ") {
			influence, length := CalcStrength(i / len(board), i % len(board), board)
			if influence > bestInfluence || influence == bestInfluence && length > bestLength {
				bestRow, bestCol, bestInfluence, bestLength = i / len(board), i % len(board), influence, length
				equalCount = 0
			} else if influence == bestInfluence && length == bestLength {
				equalCount++
				if utils.RandomNumber(0, equalCount) == 0 {
					bestRow, bestCol, bestInfluence, bestLength = i / len(board), i % len(board), influence, length
				}
			}
		}
	}
	return bestRow, bestCol
}

func FindBestMove(board [][]string) (int, int) {
	urgent, row, col := FindUrgentMove(board)
	if urgent {
		return row, col
	}
	return FindSmartestMove(board)
}

func AiTurn (mode int, board [][]string) {
	var row, col int
	switch mode {
	case ModeCpuMedium:
		row, col = FindGoodMove(board)
	case ModeCpuHard:
		row, col = FindBestMove(board)
	default:
		row, col = FindRandomMove(board)
	}
	
    PlayMove(row, col, false, board)
}
