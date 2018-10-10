package game

type Player struct {
	board *Board
	mark  string
	turns int
	name  string
}

func NewPlayer(b *Board, m string, n string) *Player {
	p := &Player{board: b, mark: m, name: n, turns: 4}
	return p
}
