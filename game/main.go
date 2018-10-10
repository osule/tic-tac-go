package game

import (
	"fmt"
)

type Runnable interface {
	Run()
}

type Game struct {
}

func (g Game) Run() {
	fmt.Println("Tic Tac Go Started!")
}

func GetGame() Runnable {
	game := Game{}
	return game
}
