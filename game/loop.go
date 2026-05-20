package game

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"slices"
	"strconv"
)

func printBoard(g *Game) {
	board := g.cells

	fmt.Printf(" %s | %s | %s \n", board[0], board[1], board[2])
	fmt.Println("-----------")
	fmt.Printf(" %s | %s | %s \n", board[3], board[4], board[5])
	fmt.Println("-----------")
	fmt.Printf(" %s | %s | %s \n", board[6], board[7], board[8])

}

func switchPlayer(g *Game) {
	if g.playerSymbol == "X" {
		g.playerSymbol = "O"
	} else {
		g.playerSymbol = "X"
	}
}

func runLoop(g *Game) {
	fmt.Println("Starting Game")
	scanner := bufio.NewScanner(os.Stdin)
	for {
		printBoard(g)

		fmt.Printf("Player %v turn\n", g.playerSymbol)
		fmt.Printf("Enter a number between 1 and 9 to place your %v or 'exit' to exit\n", g.playerSymbol)

		scanner.Scan()
		input := scanner.Text()
		if input == "exit" {
			break
		}

		positionForSymbol, err := processInput(input)
		if err != nil {
			fmt.Println(err)
			continue
		}

		err = g.place(positionForSymbol)
		if err != nil {
			fmt.Println(err)
			continue
		}

		if hasPlayerWon(g, g.playerSymbol) {
			fmt.Printf("Player %s wins!\n", g.playerSymbol)
			break
		}

		if isBoardFilled(g) {
			fmt.Println("It is a draw!")
			break
		}

		switchPlayer(g)
	}
}

func isBoardFilled(g *Game) bool {
	return !slices.Contains(g.cells, " ")
}

func hasPlayerWon(g *Game, playerSymbol string) bool {
	winningTriple := []triple{
		{0, 1, 2},
		{3, 4, 5},
		{6, 7, 8},
		{0, 3, 6},
		{1, 4, 7},
		{2, 5, 8},
		{0, 4, 8},
		{2, 4, 6},
	}
	board := g.cells

	for _, triple := range winningTriple {
		if board[triple.A] == playerSymbol && board[triple.B] == playerSymbol && board[triple.C] == playerSymbol {
			return true
		}
	}

	return false
}

func processInput(input string) (int, error) {
	integerInput, err := strconv.Atoi(input)
	if err != nil || integerInput < 1 || integerInput > 9 {
		return 0, errors.New("Please enter a number between 1 and 9")
	}
	return integerInput, nil

}
