package game

import (
	"errors"
)

func New() Game {
	cells := make([]string, 9)
	for i := range cells {
		cells[i] = " "
	}

	return Game{
		playerSymbol: "X",
		cells:        cells,
	}

}

func (g *Game) Run() {
	runLoop(g)
}

func (g *Game) place(position int) error {
	if g.cells[position-1] != " " {
		return errors.New("Position already taken. Choose another one.")
	}
	g.cells[position-1] = g.playerSymbol
	return nil
}
