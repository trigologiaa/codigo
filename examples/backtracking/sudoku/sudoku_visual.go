package sudoku

import (
	"fmt"
	"time"
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

func (sudoku *Sudoku) VisualSolve(p *widgets.Paragraph, delay time.Duration, steps *int) bool {
	for y := 0; y < 9; y++ {
		for x := 0; x < 9; x++ {
			if sudoku.board[y][x] == 0 {
				for n := 1; n <= 9; n++ {
					if sudoku.possible(x, y, n) {
						*steps++
						sudoku.board[y][x] = n
						sudoku.Render(p, x, y, *steps)
						time.Sleep(delay)
						if sudoku.VisualSolve(p, delay, steps) {
							return true
						}
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

func (sudoku *Sudoku) Render(p *widgets.Paragraph, currentX int, currentY int, steps int) {
	print := "┏━━━━━━━┳━━━━━━━┳━━━━━━━┓\n"
	for row := 0; row < 9; row++ {
		print += "┃ "
		for col := 0; col < 9; col++ {
			if col == 3 || col == 6 {
				print += "┃ "
			}
			if sudoku.board[row][col] == 0 {
				print += "[0](fg:red) "
			} else if row == currentY && col == currentX {
				print += fmt.Sprintf("[%d](fg:blue) ", sudoku.board[row][col])
			} else {
				print += fmt.Sprintf("%d ", sudoku.board[row][col])
			}
			if col == 8 {
				print += "┃"
			}
		}
		if row == 2 || row == 5 {
			print += "\n┣━━━━━━━╋━━━━━━━╋━━━━━━━┫\n"
		} else {
			print += "\n"
		}
	}
	print += "┗━━━━━━━┻━━━━━━━┻━━━━━━━┛\n"
	print += fmt.Sprintf("\n\nPasos: [%d](fg:yellow) ", steps)
	p.Text = print
	ui.Render(p)
}