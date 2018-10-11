package game

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Runnable interface {
	Run()
}

type Game struct {
}

type Reader interface {
	ReadString() string
}

type bufReader struct {
	br *bufio.Reader
}

func (rr bufReader) ReadString() string {
	v, err := rr.br.ReadString('\n')
	if err != nil {
		return ""
	}
	return strings.TrimRight(v, "\n")
}

func createPlayers(b *Board, r Reader) Players {
	ps := make([]*Player, 2)

	for i := 0; i < 2; i++ {
		fmt.Printf("Enter Player %d name: ", i)
		name := r.ReadString()

		fmt.Printf("Enter Player %d mark: ", i)
		mark := r.ReadString()

		ps[i] = NewPlayer(b, mark, name)

		fmt.Printf("Registered Player\n%s\n", ps[i])
		fmt.Println()
	}

	return Players{players: ps}
}

func switchGamePlayBetween(ps Players, rr Reader) (bool, int) {
	var r int
	var c int
	var p *Player

	i := 0
	id := 0
	for {
		id = i % 2
		p = ps.GetPlayer(id)

		if !ps.HasTurns() {
			break
		}

		fmt.Printf("Player %d (%s) turn\n", id+1, p.name)
		fmt.Println("Enter a position on the grid to place a mark e.g. 0 1")
		pos := rr.ReadString()
		fmt.Sscanf(pos, "%d %d", &r, &c)

		p.PlaceMark(r, c)

		if p.CheckGame() {
			return true, id
		}

		i++
	}
	return false, -1
}

func (g Game) Run() {
	b := NewBoard()
	rr := bufReader{br: bufio.NewReader(os.Stdin)}

	fmt.Println("Tic Tac Go Started!")
	fmt.Println(b)
	ps := createPlayers(&b, rr)
	won, _ := switchGamePlayBetween(ps, rr)
	if won {
		fmt.Println("Game ended. We have a winner!")
		fmt.Println(ps)
		fmt.Println(b)
	} else {
		fmt.Println("Game ended. We have a tie!")
	}
}

func GetGame() Runnable {
	return Game{}
}
