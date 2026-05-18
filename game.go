package main

type Game struct {
	isPlayerOnesTurn bool
	board            board
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
		isPlayerOnesTurn: true,
		board:            b,
	}

}
