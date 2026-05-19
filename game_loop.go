package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

type board struct {
	fields []string
}

func printBoard(g *Game) {
	board := g.board

	fmt.Printf(" %s | %s | %s \n", board.fields[0], board.fields[1], board.fields[2])
	fmt.Println("-----------")
	fmt.Printf(" %s | %s | %s \n", board.fields[3], board.fields[4], board.fields[5])
	fmt.Println("-----------")
	fmt.Printf(" %s | %s | %s \n", board.fields[6], board.fields[7], board.fields[8])

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

func startGameLoop(g *Game) {
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
		fmt.Printf("Your input: %v\n", positionForSymbol)
		if err != nil {
			fmt.Println(err)
			continue
		}

		changePlayer(g)
	}
}

func processInput(input string) (int, error) {
	int, err := strconv.Atoi(input)
	if err != nil || int < 1 || int > 9 {
		return 0, errors.New("Please enter a number between 1 and 9")
	}
	return int, nil

}
