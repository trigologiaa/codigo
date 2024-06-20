package main

import (
	"log"
	"time"
	"github.com/trigologiaa/codigo/examples/backtracking/sudoku"
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

func main() {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()
	p0 := widgets.NewParagraph()
	p0.SetRect(0, 0, 35, 35)
	p0.Border = false
	var board = [9][9]int{
		{4, 0, 0, 1, 9, 5, 0, 6, 8},
		{8, 0, 0, 0, 0, 0, 7, 0, 1},
		{9, 0, 1, 6, 0, 0, 0, 3, 0},
		{0, 0, 7, 0, 2, 6, 0, 0, 0},
		{5, 0, 0, 0, 0, 0, 0, 0, 3},
		{0, 1, 0, 8, 7, 0, 4, 0, 0},
		{0, 3, 0, 0, 0, 0, 8, 0, 5},
		{1, 0, 5, 0, 0, 0, 0, 0, 0},
		{7, 9, 0, 4, 0, 1, 0, 0, 0},
	}
	delay := 250 * time.Millisecond
	sudoku := sudoku.NewSudoku(board)
	steps := 0
	sudoku.VisualSolve(p0, delay, &steps)
	uiEvents := ui.PollEvents()
	for {
		e := <-uiEvents
		switch e.ID {
		case "q", "<C-c>":
			return
		}
	}
}