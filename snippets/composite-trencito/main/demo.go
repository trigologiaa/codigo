package main

import (
	"fmt"
	"github.com/trigologiaa/codigo/snippets/composite-trencito/composite"
)

func main() {
	// Defino algunas componentes básicas
	cuerpo_locomotora := composite.NewRectangulo(9, 4)
	frente_locomotora := composite.NewTriangulo(2, 4)
	cabina_locomotora := composite.NewRectangulo(2, 5)
	chimenea_1 := composite.NewRectangulo(1, 2)
	chimenea_2 := composite.NewRectangulo(1, 2)
	rueda := composite.NewCirculo(1)
	// Creo y compongo la locomotora
	locomotora := composite.NewFiguraCompuesta()
	locomotora.Add(cuerpo_locomotora)
	locomotora.Add(frente_locomotora)
	locomotora.Add(cabina_locomotora)
	locomotora.Add(chimenea_1)
	locomotora.Add(chimenea_2)
	locomotora.Add(rueda)
	locomotora.Add(rueda)
	locomotora.Add(rueda)
	locomotora.Add(rueda)
	// Creo y compongo un vagon
	cuerpo_vagon := composite.NewRectangulo(6, 5)
	vagon := composite.NewFiguraCompuesta()
	vagon.Add(cuerpo_vagon)
	vagon.Add(rueda)
	vagon.Add(rueda)
	// Creo y compongo al furgon
	furgon := composite.NewFiguraCompuesta()
	furgon.Add(composite.NewRectangulo(5, 5))
	furgon.Add(composite.NewTriangulo(2, 3))
	furgon.Add(rueda)
	furgon.Add(rueda)
	// Creo y compongo el trencito
	tren := composite.NewFiguraCompuesta()
	tren.Add(locomotora)
	tren.Add(vagon)
	tren.Add(vagon)
	tren.Add(furgon)
	// le ponemos la banderita
	banderita := composite.NewRectangulo(3, 2)
	tren.Add(banderita)
	// Para saber el area de la figura compuesta "tren"
	// solo le pedimos el Area() al tren y se hace la magia del composite
	fmt.Printf("%8.2f\n", tren.Area())
}