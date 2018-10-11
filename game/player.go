package game

import (
	"fmt"
	"strings"
)

type Player struct {
	board *Board
	mark  string
	turns int
	name  string
}

type Players struct {
	players []*Player
}

func NewPlayer(b *Board, m string, n string) *Player {
	p := &Player{board: b, mark: m, name: n, turns: 4}
	return p
}

func (p Player) HasTurns() bool {
	return p.turns > 0
}

func (p Player) CheckGame() bool {
	return p.board.CheckGame(p.mark)
}

func (p Player) PlaceMark(row int, column int) {
	if p.HasTurns() {
		p.board.PlaceMark(row, column, p.mark)
		p.turns -= 1
	}
}

func (p Player) String() string {
	return fmt.Sprintf("Name: %s\t Mark: %s \t No.turns: %d", p.name, p.mark, p.turns)
}

func (ps Players) HasTurns() bool {
	for _, p := range ps.players {
		if p.HasTurns() {
			return true
		}
	}
	return false
}

func (ps Players) GetPlayer(id int) *Player {
	return ps.players[id]
}

func (ps Players) String() string {
	b := strings.Builder{}
	for _, p := range ps.players {
		b.WriteString(p.String())
		b.WriteString("\n")
	}
	return b.String()
}
