package main

import (
	"fmt"
	"log"
	"strings"
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"github.com/trigologiaa/codigo/clase-4-ordenamiento/slice"
	"github.com/trigologiaa/codigo/clase-4-ordenamiento/sort"
)

// Parametros de la visualización. Cambiar a gusto
const tamaño = 7
const tiempo = 1000
var elementos = slice.GenerarNumerosConsecutivos(tamaño)
var p *widgets.Paragraph
var bc *widgets.BarChart

func main() {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()
	// Text
	p = widgets.NewParagraph()
	p.Text = strings.Trim(strings.Join(strings.Split(fmt.Sprint(elementos), " "), ", "), "[]")
	p.SetRect(0, 0, 100, 5)
	// Barchart
	bc = widgets.NewBarChart()
	bc.Data = elementos
	bc.SetRect(0, 5, tamaño * 2 + 2, tamaño + 5 + 3)
	bc.BarWidth = 1
	bc.BarColors = []ui.Color {ui.ColorWhite}
	bc.NumStyles = []ui.Style {ui.NewStyle(ui.ColorYellow)}
	ui.Render(bc)
	renderizar()
	uiEvents := ui.PollEvents()
	for {
		e := <-uiEvents
		switch e.ID {
		case "<C-c>":
			return
		case "1":
			p.Text = "Ordenando por inserción"
			sort.InsertionWithSleep(elementos, tiempo, renderizar, color)
			p.Text = "Esperando..."
			renderizar()
		case "2":
			p.Text = "Ordenando por selección"
			sort.SelectionWithSleep(elementos, tiempo, renderizar, color)
			p.Text = "Esperando..."
			renderizar()
		case "s":
			sort.Shuffle(elementos)
			renderizar()
		}
	}
}

func renderizar() {
	ui.Clear()
	ui.Render(p, bc)
}

func color(Colores ...int) {
	listaDeColores := []ui.Color{ui.ColorRed, ui.ColorBlue, ui.ColorYellow}
	colores := obtenerMismoColor()
	for indice, valor := range Colores {
		colores[valor] = listaDeColores[indice%len(listaDeColores)]
	}
	bc.BarColors = colores
}

func obtenerMismoColor() []ui.Color {
	colores := make([]ui.Color, tamaño)
	for indice := range colores {
		colores[indice] = ui.ColorGreen
	}
	return colores
}