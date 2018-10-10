package game

import (
	"strings"
)

type Board struct {
	grid [][]string
}

// Create a board
func (Board) Board() {
}

// Check if there are matching cells in a direction
func (Board) CheckGame(mark string) bool {
	return true
}

// Place mark in location on board
func (Board) PlaceMark(row int, column int, mark string) {
}

// Implement Stringer interface
func (Board) String() string {
	return strings.Repeat("- - - -\n", 4)
}
