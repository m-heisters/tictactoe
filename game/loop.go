package game

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

func printBoard(g *Game) {
	board := g.board

	fmt.Printf(" %s | %s | %s \n", board.cells[0], board.cells[1], board.cells[2])
	fmt.Println("-----------")
	fmt.Printf(" %s | %s | %s \n", board.cells[3], board.cells[4], board.cells[5])
	fmt.Println("-----------")
	fmt.Printf(" %s | %s | %s \n", board.cells[6], board.cells[7], board.cells[8])

}

func changePlayer(g *Game) error {
	switch g.playerSymbol {
	case "X":
		g.playerSymbol = "O"
		return nil
	case "O":
		g.playerSymbol = "X"
		return nil
	default:
		return errors.New("Unknown player symbol. Exiting game")
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

		err = g.DrawSymbol(positionForSymbol)
		if err != nil {
			fmt.Println(err)
			continue
		}

		isFinished := doesPlayerWin(g)
		if isFinished {
			fmt.Printf("Player %s wins!\n", g.playerSymbol)
			break
		}

		changePlayer(g)
	}
}

func doesPlayerWin(g *Game) bool {
	winningTriple := []Triple{
		{0, 1, 2},
		{3, 4, 5},
		{6, 7, 8},
		{0, 3, 6},
		{1, 4, 7},
		{2, 5, 8},
		{0, 4, 8},
		{2, 4, 6},
	}

	for _, triple := range winningTriple {
		cells := g.board.cells
		// fmt.Printf("triple.A: %v, triple.B: %v, triple.C: %v\n", cells[triple.A], cells[triple.B], cells[triple.C])
		if cells[triple.A] == g.playerSymbol && cells[triple.B] == g.playerSymbol && cells[triple.C] == g.playerSymbol {
			return true
		}
	}

	return false
}

func processInput(input string) (int, error) {
	int, err := strconv.Atoi(input)
	if err != nil || int < 1 || int > 9 {
		return 0, errors.New("Please enter a number between 1 and 9")
	}
	return int, nil

}
