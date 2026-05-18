package main

import (
	"bufio"
	"fmt"
	"os"
)

type board struct {
	fields []string
}

func printBoard(g Game) {
	board := g.board

	fmt.Printf(" %s | %s | %s \n", board.fields[0], board.fields[1], board.fields[2])
	fmt.Println("-----------")
	fmt.Printf(" %s | %s | %s \n", board.fields[3], board.fields[4], board.fields[5])
	fmt.Println("-----------")
	fmt.Printf(" %s | %s | %s \n", board.fields[6], board.fields[7], board.fields[8])

}

func promptInput(g Game) {
	var playerSymbol string
	if g.isPlayerOnesTurn {
		playerSymbol = "X"
	} else {
		playerSymbol = "O"
	}

	fmt.Printf("Player %v turn\n", playerSymbol)

}

func startGameLoop(g Game) {
	fmt.Println("Starting Game")
	scanner := bufio.NewScanner(os.Stdin)
	for {

		printBoard(g)
		promptInput(g)

		scanner.Scan()

		input := scanner.Text()

		if input == "exit" {
			break
		}
		g.isPlayerOnesTurn = !g.isPlayerOnesTurn

	}
}
