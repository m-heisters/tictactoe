package main

import (
	"fmt"
)

func main() {

	board := make([]string, 9)
	for i := range board {
		board[i] = " "
	}

	fmt.Printf(" %s | %s | %s \n", board[0], board[1], board[2])
	fmt.Println("----------")
	fmt.Printf(" %s | %s | %s \n", board[3], board[4], board[5])
	fmt.Println("----------")
	fmt.Printf(" %s | %s | %s \n", board[6], board[7], board[8])
}
