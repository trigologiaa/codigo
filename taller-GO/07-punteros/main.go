package main

import (
	"fmt"
	"taller-GO/07-punteros/punteros"
)

func main() {
	var x int = 10
	y := 10
	p1 := &x
	var p2 *int
	p2 = p1
	x++
	*p1++
	fmt.Println("x = ", *p2)
	fmt.Println("p2 = ", p2)
	p2 = &y
	fmt.Println("y = ", *p2)
	fmt.Println("p2 = ", p2)
	// & operador de indirección
	punteros.Incrementar(&x)
	fmt.Println("x = ", x)
	punteros.Incrementar2(x)
	fmt.Println("x = ", x)
}