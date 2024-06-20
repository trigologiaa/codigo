package ratmaze

import (
	"fmt"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestLaberinto(t *testing.T) {
	O := true
	X := false
	var maze = [][]bool{
		{O, O, X, O, O, O, O},
		{O, O, X, O, X, X, O},
		{O, O, X, O, O, X, O},
		{O, O, X, X, O, X, O},
		{O, O, O, O, O, X, O},
	}
	ratmaze := NewRatMaze(maze)
	assert.True(t, ratmaze.Solve(), fmt.Sprint("No fue dada una solución"))
	ratmaze.PrintSolution()
	assert.True(t, isRatMazeSolutionValid(ratmaze.solution, 0, 0, 6, 4), fmt.Sprint("La solución es invalida"))
	assert.True(t, isSameMaze(ratmaze.maze, ratmaze.solution), fmt.Sprint("La solución no corresponde con la maze"))
}

func TestVacio(t *testing.T) {
	O := true
	// X := false
	var maze = [][]bool{
		{O, O, O, O, O, O, O, O, O, O, O, O},
		{O, O, O, O, O, O, O, O, O, O, O, O},
		{O, O, O, O, O, O, O, O, O, O, O, O},
		{O, O, O, O, O, O, O, O, O, O, O, O},
		{O, O, O, O, O, O, O, O, O, O, O, O},
		{O, O, O, O, O, O, O, O, O, O, O, O},
		{O, O, O, O, O, O, O, O, O, O, O, O},
		{O, O, O, O, O, O, O, O, O, O, O, O},
		{O, O, O, O, O, O, O, O, O, O, O, O},
		{O, O, O, O, O, O, O, O, O, O, O, O},
		{O, O, O, O, O, O, O, O, O, O, O, O},
		{O, O, O, O, O, O, O, O, O, O, O, O},
	}
	ratmaze := NewRatMaze(maze)
	assert.True(t, ratmaze.Solve(), fmt.Sprint("No fue dada una solución"))
	ratmaze.PrintSolution()
	assert.True(t, isRatMazeSolutionValid(ratmaze.solution, 0, 0, 11, 11), fmt.Sprint("La solución es invalida"))
	assert.True(t, isSameMaze(ratmaze.maze, ratmaze.solution), fmt.Sprint("La solución no corresponde con la maze"))
}

func isSameMaze(original [][]bool, solved [][]bool) bool {
	// Recorro todo el tablero
	for y := 0; y < len(original); y++ {
		for x := 0; x < len(original[y]); x++ {
			if !original[y][x] && solved[y][x] {
				return false
			}
		}
	}
	return true
}

func isRatMazeSolutionValid(maze [][]bool, origX int, origY int, destX int, destY int) bool {
	// Compruebo origen y destino
	if !maze[origY][origX] || !maze[destY][destX] {
		return false
	}
	// Compruebo adyacencias
	rows := len(maze)
	for y := 0; y < rows; y++ {
		cols := len(maze[y])
		for x := 0; x < len(maze[y]); x++ {
			// Si soy parte de solución, compruebo tener al menos 1 adyacente
			if maze[y][x] &&
				(y != 0 && !maze[y-1][x]) &&
				(y != rows-1 && !maze[y+1][x]) &&
				(x != 0 && !maze[y][x-1]) &&
				(y != cols-1 && !maze[y][x+1]) {
				fmt.Print(x, y, cols, rows)
				return false
			}
		}
	}
	return true
}