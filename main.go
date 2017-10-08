package main

import "fmt"
import "bufio"
import "strconv"
import "os"

func printBoard(board string) {
	fmt.Println("=========")
	fmt.Println("||", board[0:3], "||")
	fmt.Println("||", board[3:6], "||")
	fmt.Println("||", board[6:9], "||")
	fmt.Println("=========")
}

func winner(board string, player int) bool {
	var newMarker byte

	if player == 1 {
		newMarker = 'X'
	} else {
		newMarker = 'O'
	}

	winner := (board[0] == board[1]) && (board[1] == board[2]) && (board[0] == newMarker) || // top horiz
		(board[3] == board[4]) && (board[4] == board[5]) && (board[3] == newMarker) || // middle horiz
		(board[6] == board[7]) && (board[7] == board[8]) && (board[7] == newMarker) || // bottom horiz
		(board[0] == board[3]) && (board[0] == board[6]) && (board[0] == newMarker) || // left vert
		(board[1] == board[4]) && (board[4] == board[7]) && (board[1] == newMarker) || // middle vert
		(board[2] == board[5]) && (board[2] == board[8]) && (board[2] == newMarker) || // right vert
		(board[0] == board[4]) && (board[0] == board[8]) && (board[0] == newMarker) || // diag 1
		(board[6] == board[4]) && (board[6] == board[2]) && (board[2] == newMarker) // diag 2

	return winner

}

func moveValid(board, move string) bool {
	mv, err := strconv.Atoi(move)
	mv -= 1

	if (len(board) <= mv) || mv < 0 {
		return false
	}

	currentValue := board[mv]

	if (currentValue == 'X') || (currentValue == 'O') || err != nil {
		return false
	} else {
		return true
	}
}

func updateBoard(board, move string, activePlayer int) string {
	mv, _ := strconv.Atoi(move)
	mv -= 1

	out := []rune(board)

	var newMarker rune

	if activePlayer == 1 {
		newMarker = 'X'
	} else {
		newMarker = 'O'
	}

	out[mv] = newMarker
	return string(out)
}

func isDraw(board string) bool {
	for _, j := range board {
		if (j != 'X') && (j != 'O') {
			return false
		}
	}
	return true
}

func main() {
	board := "123456789"
	activePlayer := 1

	// main loop
	for true {
		printBoard(board)
		fmt.Printf("Next player: Player %d\n", activePlayer)
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter Move: ")
		move, _ := reader.ReadString('\n')
		move = move[:len(move)-1]

		if moveValid(board, move) {
			board = updateBoard(board, move, activePlayer)

			if winner(board, activePlayer) {
				printBoard(board)
				fmt.Printf("We have a winner!! Congratulations Player %d\n", activePlayer)
				os.Exit(0)
			}
			if isDraw(board) {
				fmt.Println("We have a draw!!")
				os.Exit(0)
			}

			activePlayer = activePlayer%2 + 1

		} else {
			fmt.Println("Invalid move, please try again")
		}

	}

}
