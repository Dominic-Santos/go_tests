package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"math/rand"
	"time"
)

var reader = bufio.NewReader(os.Stdin)

var mode_cpu int = 1
var mode_2play int = 2

func randomNumber(min int, max int) int{
    rand.Seed(time.Now().UnixNano())
    return rand.Intn(max - min + 1) + min
}

func readUserInput() string {
	text, _ := reader.ReadString('\n')
	return text[:len(text) - 1]
}

func requestUserInput() string {
	fmt.Print("->")
	return readUserInput()
}

func requestUserInt() (int, error) {
	return strconv.Atoi(requestUserInput())
}

func showHeader(){
	fmt.Println("Tic Tac Toe - GO - by Dominic Santos\n")
}

func showError(s string){
	fmt.Printf("Invalid %s - please try again\n", s)
}

func showTitleScreen(error bool) {
	showHeader()
	fmt.Println("Please select a mode:\n%v - Vs AI\n%v - 2 Player\n\n0 - Quit\n", mode_cpu, mode_2play)
	if error {
		showError("mode")
	}
}

func title() int {
	error := false
	for {
		clear()
		showTitleScreen(error)
		n, err := requestUserInt()
		error = err != nil || n < 0 || n > 2 

		if !error {
			return n
		}
	}
}

func clear() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func showBoardSize(error bool) {
	showHeader()
	fmt.Println("Please select a board size (min: 3, max: 9), 0 = back:\n")
	if error {
		showError("board size")
	}
}

func boardSize() int {
	error := false
	for {
		clear()
		showBoardSize(error)
		n, err := requestUserInt()
		error = err != nil || n > 9 || n != 0 && n < 3

		if !error {
			return n
		}
	}
}

func makeBoard(size int) [][]string {
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

func intToChar(i int) string {
	return string('A' - 1 + i)
}

func charToInt(s string) int {
	return int([]rune(s)[0])
}

func showBoard(board [][]string) {
	fmt.Print("  ")
	for i := 1; i <= len(board); i++ {
       fmt.Printf("%v   ", i)
	}
	fmt.Print("\n")
	for i := 0; i < len(board); i++ {
		if i != 0 {
			fmt.Printf(" %s%s\n", strings.Repeat("---+", len(board) - 1), "---")
		}
		fmt.Printf("%s %s \n", intToChar(i + 1), strings.Join(board[i], " | "))
	}
	fmt.Println("")
}

func showPrompt(mode int, turn bool) {
	var player string
	if mode == mode_2play {
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

func showGame(board [][]string, mode int, turn bool, err bool) {
	clear()
	showHeader()
	showBoard(board)
	showPrompt(mode, turn)
	if err {
		showError("move")
	}
}

func validMove(row int, col int, board [][]string) bool {
	return row < len(board) && col < len(board) && board[row][col] == " "
}

func readMove(move string) (int, int, bool) {
	if len(move) >= 2 {
		row := charToInt(move[0:1])
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

func playMove(row, col int, turn bool, board [][]string) {
	if turn {
		board[row][col] = "O"
	} else {
		board[row][col] = "X"
	}
}

func playGame(mode int, size int) {
	board := makeBoard(size)
	turn := randomNumber(0, 1) == 0
	err := false
	for {
		showGame(board, mode, turn, err)
		err = true
		row, col, e := readMove(requestUserInput())
		if !e {
			if validMove(row, col, board){
				err = false
				playMove(row, col, turn, board)
			}
		}
	}
}

func main() {
	playGame(1, 3)
	return
	mainloop:
	for {
		switch mode := title(); mode {
		case 0:
			break mainloop
		default:
			size := boardSize()
			if size != 0 {
				playGame(mode, size)
			}
		}
	}
}

