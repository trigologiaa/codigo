package main

import (
	"fmt"
	"log"
	"reflect"
	"time"
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

var renderizable = []ui.Drawable{}
const ancho = 14
const alto = 6

func principal() {
	var i int = 42
	var j int = 2701
	p := &i
	p2 := p
	// Agrego para mostrar
	agregarVariable("i", &i, 0, 0)
	agregarVariable("j", &j, 1, 0)
	agregarVariableP("p", &p, 2, 0)
	agregarVariableP("p2", &p2, 3, 0)
	actualizarConsola()
	time.Sleep(5 * time.Second)
	*p = *p * 2
	agregarVariable("i", &i, 0, 1)
	agregarVariable("j", &j, 1, 1)
	agregarVariableP("p", &p, 2, 1)
	agregarVariableP("p2", &p2, 3, 1)
	actualizarConsola()
}

func main() {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()
	// Codigo
	principal()
	actualizarConsola()
	uiEvents := ui.PollEvents()
	for {
		e := <-uiEvents
		switch e.ID {
		case "q", "<C-c>":
			return
		}
	}
}

func agregarVariable(nombre string, direccion *int, columna int, fila int) {
	pi := widgets.NewParagraph()
	pi.Text = fmt.Sprint(nombre, " [", reflect.TypeOf(*direccion), "](fg:blue)", "\n", "[", *direccion, "](fg:red,mod:bold)", "\n", direccion)
	pi.SetRect(columna*ancho, fila*alto, (columna+1)*ancho, (fila+1)*alto)
	//pi.Border = false
	renderizable = append(renderizable, pi)
}

func agregarVariableP(nombre string, direccion **int, columna int, fila int) {
	pi := widgets.NewParagraph()
	pi.Text = fmt.Sprint(nombre, " [", reflect.TypeOf(*direccion), "](fg:blue)", "\n", "[", *direccion, "](fg:red,mod:bold)", "\n", direccion, "\n", "[", **direccion, "](fg:yellow,mod:bold)")
	pi.SetRect(columna*ancho, fila*alto, (columna+1)*ancho, (fila+1)*alto)
	//pi.Border = false
	renderizable = append(renderizable, pi)
}

func actualizarConsola() {
	ui.Render(renderizable...)
}