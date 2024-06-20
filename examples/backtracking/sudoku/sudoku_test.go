package sudoku

import (
	"fmt"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestSudokuUnicaSolucion(t *testing.T) {
	var board = [9][9]int{
		{0, 0, 0, 1, 0, 5, 0, 6, 8},
		{0, 0, 0, 0, 0, 0, 7, 0, 1},
		{9, 0, 1, 0, 0, 0, 0, 3, 0},
		{0, 0, 7, 0, 2, 6, 0, 0, 0},
		{5, 0, 0, 0, 0, 0, 0, 0, 3},
		{0, 0, 0, 8, 7, 0, 4, 0, 0},
		{0, 3, 0, 0, 0, 0, 8, 0, 5},
		{1, 0, 5, 0, 0, 0, 0, 0, 0},
		{7, 9, 0, 4, 0, 1, 0, 0, 0},
	}
	sudoku := NewSudoku(board)
	sudoku.Solve()
	// sudoku.PrintSolutions()
	assert.Len(t, sudoku.solutions, 1, "La solución debe ser una")
	for i := range sudoku.solutions {
		assert.True(t, isSudokuValid(sudoku.solutions[i]), fmt.Sprint("La solución ", i+1, " del Sudoku es invalida"))
		assert.True(t, isSameSudoku(board, sudoku.solutions[i]), fmt.Sprint("La solución ", i+1, " del Sudoku no corresponde con la original"))
	}
}

func TestSudokuMultipleSolucion(t *testing.T) {
	var board = [9][9]int{
		{5, 3, 0, 0, 7, 0, 0, 0, 0},
		{6, 0, 0, 1, 9, 5, 0, 0, 0},
		{0, 9, 8, 0, 0, 0, 0, 6, 0},
		{8, 0, 0, 0, 6, 0, 0, 0, 3},
		{4, 0, 0, 8, 0, 3, 0, 0, 1},
		{7, 0, 0, 0, 2, 0, 0, 0, 6},
		{0, 6, 0, 0, 0, 0, 2, 8, 0},
		{0, 0, 0, 4, 1, 9, 0, 0, 5},
		{0, 0, 0, 0, 8, 0, 0, 0, 0},
	}
	sudoku := NewSudoku(board)
	sudoku.Solve()
	// sudoku.PrintSolutions()
	assert.Len(t, sudoku.solutions, 2, "La soluciones deben ser dos")
	for i := range sudoku.solutions {
		assert.True(t, isSudokuValid(sudoku.solutions[i]), fmt.Sprint("La solución ", i+1, " del Sudoku es invalida"))
		assert.True(t, isSameSudoku(board, sudoku.solutions[i]), fmt.Sprint("La solución ", i+1, " del Sudoku no corresponde con la original"))
	}
}

func TestSudokuMilesSoluciones(t *testing.T) {
	var board = [9][9]int{
		{1, 2, 3, 4, 5, 6, 7, 8, 9},
		{4, 5, 6, 7, 8, 9, 1, 2, 3},
		{7, 8, 9, 0, 0, 0, 4, 5, 6},
		{2, 3, 1, 0, 0, 0, 0, 0, 0},
		{5, 6, 4, 0, 0, 0, 0, 0, 0},
		{8, 9, 7, 0, 0, 0, 0, 0, 0},
		{3, 1, 2, 0, 0, 0, 0, 0, 0},
		{6, 4, 5, 0, 0, 0, 0, 0, 0},
		{9, 7, 8, 0, 0, 0, 0, 0, 0},
	}
	sudoku := NewSudoku(board)
	sudoku.Solve()
	assert.Len(t, sudoku.solutions, 32604, "La soluciones deben ser 32604")
	for i := range sudoku.solutions {
		assert.True(t, isSudokuValid(sudoku.solutions[i]), fmt.Sprint("La solución ", i+1, " del Sudoku es invalida"))
		assert.True(t, isSameSudoku(board, sudoku.solutions[i]), fmt.Sprint("La solución ", i+1, " del Sudoku no corresponde con la original"))
	}
}

func isSameSudoku(original [9][9]int, solved [9][9]int) bool {
	// Recorro todo el tablero
	for y := 0; y < 9; y++ {
		for x := 0; x < 9; x++ {
			if original[y][x] != 0 && original[y][x] != solved[y][x] {
				return false
			}
		}
	}
	return true
}

func isSudokuValid(board [9][9]int) bool {
	// Recorro todo el tablero
	for y := 0; y < 9; y++ {
		for x := 0; x < 9; x++ {
			n := board[y][x]
			// Analizo EL RESTO la columna en busqueda de repetidos
			for i := x + 1; i < 9; i++ {
				if board[y][i] == n {
					return false
				}
			}
			// Analizo EL RESTO de la fila en busqueda de repetidos
			for i := y + 1; i < 9; i++ {
				if board[i][x] == n && y != i {
					return false
				}
			}
			// Analizo el cuadrante en busqueda de repetidos
			y0 := (y / 3) * 3
			x0 := (x / 3) * 3
			// TODO: Mejorar con 1 solo ciclo de 9, y optimizar para no analizar duplicado
			for i := 0; i < 3; i++ {
				for j := 0; j < 3; j++ {
					if board[y0+i][x0+j] == n && y != y0+i && x != x0+j {
						return false
					}
				}
			}
		}
	}
	return true
}