package ratmaze

import (
	"fmt"
)

type RatMaze struct {
	maze     [][]bool
	solution [][]bool
	destX    int
	destY    int
}

func NewRatMaze(maze [][]bool) *RatMaze {
	return &RatMaze{maze: maze}
}

// Controla que la posición esté dentro del laberinto, y que la posición sea un camino despejado
func (ratmaze *RatMaze) isSafe(x int, y int) bool {
	return y >= 0 && y < len(ratmaze.maze) && x >= 0 && x < len(ratmaze.maze[y]) && ratmaze.maze[y][x]
}

func (ratmaze *RatMaze) Solve() bool {
	// Preparamos la solución creando una matriz del mismo tamaño lleno de falsos
	rows := len(ratmaze.maze)
	ratmaze.solution = make([][]bool, rows)
	for i := 0; i < rows; i++ {
		cols := len(ratmaze.maze[i])
		ratmaze.solution[i] = make([]bool, cols)
		// for j := 0; j < cols; j++ {
		// 	ratmaze.solution[i][j] = false
		// }
	}
	// El destino es la posición de más abajo a la derecha
	ratmaze.destY = rows - 1
	ratmaze.destX = len(ratmaze.maze[rows-1]) - 1
	// Comienza en el 0, 0
	return ratmaze.solve(0, 0)
}

func (ratmaze *RatMaze) solve(x int, y int) bool {
	// Si estoy en el destino, devuelvo verdadero
	if x == ratmaze.destX && y == ratmaze.destY && ratmaze.maze[y][x] {
		ratmaze.solution[y][x] = true
		return true
	}
	// Si la posición no es válida
	if !ratmaze.isSafe(x, y) {
		return false
	}
	// Si estoy en una posición que ya es solución
	if ratmaze.solution[y][x] {
		return false
	}
	ratmaze.solution[y][x] = true
	// Mostrar pasos
	// ratmaze.PrintSolution()
	// fmt.Println()
	// Busco para derecha
	if ratmaze.solve(x+1, y) {
		return true
	}
	// Busco para abajo
	if ratmaze.solve(x, y+1) {
		return true
	}
	// Busco para izquierda
	if ratmaze.solve(x-1, y) {
		return true
	}
	// Busco para arriba
	if ratmaze.solve(x, y-1) {
		return true
	}
	// Si nada fue valido, backtrackeo dejando la posición actual como no parte de la solución
	ratmaze.solution[y][x] = false
	return false
}

func (ratmaze *RatMaze) PrintSolution() {
	for row := 0; row < len(ratmaze.solution); row++ {
		for col := 0; col < len(ratmaze.solution[row]); col++ {
			x := " "
			if ratmaze.solution[row][col] {
				x = "X"
			}
			fmt.Print(x)
		}
		fmt.Println()
	}
}