package main

import (
	"fmt"
	"strconv"
	"strings"
	"os"
	"os/exec"
	"bufio"
)

import "github.com/Dominic-Santos/gotest/utils"

var reader = bufio.NewReader(os.Stdin)

func Clear() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func RequestUserInput() string {
	fmt.Print("->")
	return ReadUserInput()
}

func ReadUserInput() string {
	text, _ := reader.ReadString('\n')
	return text[:len(text) - 1]
}

func RequestUserInt() (int, error) {
	return strconv.Atoi(RequestUserInput())
}

func ShowHeader(){
	fmt.Println("Tic Tac Toe - GO - by Dominic Santos\n")
}

func ShowError(s string){
	fmt.Printf("Invalid %s - please try again\n", s)
}

func ShowTitleScreen(error bool) {
	ShowHeader()
	fmt.Printf("Please select a mode:\n%v - Vs AI\n%v - 2 Player\n\n0 - Quit\n", ModeCpu, Mode2Play)
	if error {
		ShowError("mode")
	}
}

func ShowBoardSize(error bool) {
	ShowHeader()
	fmt.Println("Please select a board size (min: 3, max: 9), 0 = back:\n")
	if error {
		ShowError("board size")
	}
}

func ShowBoard(board [][]string) {
	fmt.Print("  ")
	for i := 1; i <= len(board); i++ {
       fmt.Printf("%v   ", i)
	}
	fmt.Print("\n")
	for i := 0; i < len(board); i++ {
		if i != 0 {
			fmt.Printf(" %s%s\n", strings.Repeat("---+", len(board) - 1), "---")
		}
		fmt.Printf("%s %s \n", utils.IntToChar(i + 1), strings.Join(board[i], " | "))
	}
	fmt.Println("")
}

func ShowPrompt(mode int, turn bool) {
	var player string
	if mode == Mode2Play {
		if turn {
			player = "O's"
		} else {
			player = "X's"
		}
	} else {
		player = "your"
	}

	fmt.Printf("It's %s turn, please select a coordenate to play (example: A1)\n\n", player)

}

func ShowGame(board [][]string, mode int, turn bool, err bool) {
	Clear()
	ShowHeader()
	ShowBoard(board)
	ShowPrompt(mode, turn)
	if err {
		ShowError("move")
	}
}

func ReadMove(move string) (int, int, bool) {
	if len(move) >= 2 {
		row := utils.CharToInt(move[0:1])
		if row >= 97 {
			row -= 97
		} else {
			row -= 65
		}
		col, err := strconv.Atoi(move[1:2])
		if err == nil {
			return row, col - 1, false
		} 
	}
	return 0, 0, true
}

func ShowResult(state string, mode int) {
	if state == "D" {
		fmt.Println("Game is a Draw!\n")
	} else if mode == Mode2Play {
		fmt.Printf("Player %s Wins!\n\n", state)
	} else if state == "O" {
		fmt.Println("You Win!\n")
	} else {
		fmt.Println("You Lose!\n")
	}
}

func ShowGameResult(board [][]string, mode int, state string) {
	Clear()
	ShowHeader()
	ShowBoard(board)
	ShowResult(state, mode)
}
