package main

var Debug = false

func Title() int {
	error := false
	for {
		Clear()
		ShowTitleScreen(error)
		n, err := RequestUserInt()
		error = err != nil || n < 0 || n > 4 

		if !error {
			return n
		}
	}
}

func BoardSize() int {
	error := false
	for {
		Clear()
		ShowBoardSize(error)
		n, err := RequestUserInt()
		error = err != nil || n > 9 || n != 0 && n < 3

		if !error {
			return n
		}
	}
}

func main() {
	mainloop:
	for {
		switch mode := Title(); mode {
		case 0:
			break mainloop
		default:
			size := BoardSize()
			if size != 0 {
				PlayGame(mode, size)
			}
		}
	}
}

