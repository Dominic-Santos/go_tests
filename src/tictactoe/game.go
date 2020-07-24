package main

import "strings"
import "github.com/Dominic-Santos/gotest/utils"

var ModeCpu int = 1
var Mode2Play int = 2

func NumberAvailableMoves(board [][]string) int {
	count := 0
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board); j++ {
			if board[i][j] == " " {
				count++
			}
		} 
	}
	return count
}

func ValidMove(row int, col int, board [][]string) bool {
	return row < len(board) && col < len(board) && board[row][col] == " "
}

func PlayMove(row, col int, turn bool, board [][]string) {
	if turn {
		board[row][col] = "O"
	} else {
		board[row][col] = "X"
	}
}

func MakeBoard(size int) [][]string {
	board := make([][]string, size, size)
	
	for i := 0; i < size; i++ {
		row := make([]string, size, size)
		for j := 0; j < size; j++ {
			row[j] = " "
		}
		board[i] = row
	}

	return board
}

func CheckBoard(board [][]string) string {
	// check rows win
	for row := 0; row < len(board); row++ {
		if board[row][0] != " " && strings.Join(board[row], "") == strings.Repeat(board[row][0], len(board)) {
			return board[row][0]
		}
	}

	// check columns win
	for col := 0; col < len(board); col++ {
		if board[0][col] != " " {
			var cur []string
			for row := 0; row < len(board); row++ {
				cur = append(cur, board[row][col])
			}
			if strings.Join(cur, "") == strings.Repeat(board[0][col], len(board)) {
				return board[0][col]
			}
		}
	}
	// check diagonal win
	

	// check board full
	if NumberAvailableMoves(board) == 0 {
		return "D"
	}

	return "G"
}

func PlayGame(mode int, size int) {
	board := MakeBoard(size)
	turn := utils.RandomNumber(0, 1) == 0
	err := false
	game := true
	for game {
		if !turn && mode != Mode2Play {
			AiTurn(mode, board)
			err = false
		}

		if turn || mode == Mode2Play {
			ShowGame(board, mode, turn, err)
			err = true
			
			row, col, e := ReadMove(RequestUserInput())
			if !e {
				if ValidMove(row, col, board){
					err = false
					PlayMove(row, col, turn, board)
				}
			}
		}

		if !err {
			state := CheckBoard(board)
			if state != "G" {
				game = false
				ShowGameResult(board, mode, state)
				RequestUserInput()
			}
			turn = !turn
		}
	}
}