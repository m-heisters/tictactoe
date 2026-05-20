package game

type Game struct {
	playerSymbol string
	board        board
}

type board struct {
	cells []string
}

type Triple struct {
	A, B, C int
}
