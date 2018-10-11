package game

import (
	"strings"
)

type Board struct {
	grid [][]string
}

const GRID_SIZE int = 4

// Create a board
func NewBoard() Board {
	g := make([][]string, GRID_SIZE)
	b := Board{grid: g}
	for i := 0; i < GRID_SIZE; i++ {
		b.grid[i] = make([]string, GRID_SIZE)
		for j := 0; j < GRID_SIZE; j++ {
			b.grid[i][j] = "-"
		}
	}
	return b
}

func (b *Board) matchAlongRow(mark string) bool {
	for i := 0; i < GRID_SIZE; i++ {
		matchFound := true
		for j := 0; j < GRID_SIZE; j++ {
			matchFound = matchFound && b.grid[i][j] == mark
		}
		if matchFound {
			return true
		}
	}
	return false
}

func (b *Board) matchAlongDiagonal(mark string) bool {
	matchFound := true
	for i := 0; i < GRID_SIZE; i++ {
		matchFound = matchFound && b.grid[i][i] == mark
	}
	if matchFound {
		return true
	}
	return false
}

func (b *Board) matchAlongColumn(mark string) bool {
	for i := 0; i < GRID_SIZE; i++ {
		matchFound := true
		for j := 0; j < GRID_SIZE; j++ {
			matchFound = matchFound && b.grid[j][i] == mark
		}
		if matchFound {
			return true
		}
	}
	return false
}

// Check if there are matching cells in a direction
func (b *Board) CheckGame(mark string) bool {
	return b.matchAlongRow(mark) || b.matchAlongDiagonal(mark) || b.matchAlongColumn(mark)
}

// Place mark in location on board
func (b *Board) PlaceMark(row int, column int, mark string) {
	b.grid[row][column] = mark
}

// Implement Stringer interface
func (b Board) String() string {
	builder := strings.Builder{}

	for i := 0; i < GRID_SIZE; i++ {
		builder.WriteString(strings.Join(b.grid[i], " "))
		builder.Write([]byte("\n"))
	}
	return builder.String()
}
