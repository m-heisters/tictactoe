package game

import (
	"errors"
)

func New() Game {
	cells := make([]string, 9)
	for i := range cells {
		cells[i] = " "
	}

	b := cells
	return Game{
		playerSymbol: "X",
		board:        b,
	}

}

func (g *Game) Run() {
	runLoop(g)
}

func (g *Game) DrawSymbol(position int) error {
	if g.board[position-1] != " " {
		return errors.New("Position already taken. Choose another one.")
	}
	g.board[position-1] = g.playerSymbol
	return nil
}
