package main

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
