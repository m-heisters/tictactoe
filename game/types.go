package game

type Game struct {
	playerSymbol string
	board        board
}

type board struct {
	fields []string
}
