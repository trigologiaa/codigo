package main

import (
	"fmt"
	"time"
	"taller-GO/05-condicionales/condicionales"
)

func main() {
	var numero int
	fmt.Print("Ingrese un entero: ")
	fmt.Scanln(&numero)
	condicionales.Condicional(numero)
	condicionales.SwtichBasico()
	condicionales.SwitchSinCondicion(time.Now().Local().Hour())
	condicionales.SwitchMultiple(' ')
	condicionales.SwitchFallthrough(2)
}