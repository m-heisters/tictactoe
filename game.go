package main

import (
	"errors"
)

type Game struct {
	playerSymbol string
	board        board
}

func NewGame() Game {
	fields := make([]string, 9)
	for i := range fields {
		fields[i] = " "
	}

	b := board{
		fields: fields,
	}

	return Game{
		playerSymbol: "X",
		board:        b,
	}

}

func (g *Game) DrawSymbol(position int) error {
	if g.board.fields[position-1] != " " {
		return errors.New("Position already taken. Choose another one.")
	}
	g.board.fields[position-1] = g.playerSymbol
	return nil
}
