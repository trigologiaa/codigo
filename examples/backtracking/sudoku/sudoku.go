package sudoku

import (
	"fmt"
)

type Sudoku struct {
	board     [9][9]int
	solutions [][9][9]int
}

func NewSudoku(board [9][9]int) *Sudoku {
	return &Sudoku{board: board}
}

func (sudoku *Sudoku) possible(x int, y int, n int) bool {
	// Analizo la columna en busqueda de repetidos
	for i := 0; i < 9; i++ {
		if sudoku.board[y][i] == n {
			return false
		}
	}
	// Analizo la fila en busqueda de repetidos
	for i := 0; i < 9; i++ {
		if sudoku.board[i][x] == n {
			return false
		}
	}
	// Analizo el cuadrante en busqueda de repetidos
	// Inicio del cuadrante a analizar
	y0 := (y / 3) * 3 // Division entera
	x0 := (x / 3) * 3 // Division entera
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if sudoku.board[y0+i][x0+j] == n {
				return false
			}
		}
	}
	return true
}

func (sudoku *Sudoku) Solve() bool {
	// Crea o reinicia las soluciones
	sudoku.solutions = [][9][9]int{}
	return sudoku.solve()
}

func (sudoku *Sudoku) solve() bool {
	for y := 0; y < 9; y++ {
		for x := 0; x < 9; x++ {
			if sudoku.board[y][x] == 0 {
				for n := 1; n <= 9; n++ {
					if sudoku.possible(x, y, n) {
						sudoku.board[y][x] = n
						sudoku.solve() // if true, tengo al menos 1 solucion y puedo salir
						sudoku.board[y][x] = 0
					}
				}
				return false
			}
		}
	}
	// Solucion correcta encontrada
	sudoku.addSolution()
	return true
}

func (sudoku *Sudoku) addSolution() {
	duplicate := [9][9]int{}
	for y := 0; y < 9; y++ {
		for x := 0; x < 9; x++ {
			duplicate[y][x] = sudoku.board[y][x]
		}
	}
	sudoku.solutions = append(sudoku.solutions, duplicate)
}

func (sudoku *Sudoku) PrintBoard() {
	sudoku.print(sudoku.board)
}

func (sudoku *Sudoku) PrintSolutions() {
	for i := range sudoku.solutions {
		sudoku.print(sudoku.solutions[i])
	}
}

func (sudoku *Sudoku) print(board [9][9]int) {
	fmt.Println("┏━━━━━━━┳━━━━━━━┳━━━━━━━┓")
	for row := 0; row < 9; row++ {
		fmt.Print("┃ ")
		for col := 0; col < 9; col++ {
			if col == 3 || col == 6 {
				fmt.Print("┃ ")
			}
			fmt.Printf("%d ", board[row][col])
			if col == 8 {
				fmt.Print("┃")
			}
		}
		if row == 2 || row == 5 {
			fmt.Println("\n┣━━━━━━━╋━━━━━━━╋━━━━━━━┫")
		} else {
			fmt.Println()
		}
	}
	fmt.Println("┗━━━━━━━┻━━━━━━━┻━━━━━━━┛")
}