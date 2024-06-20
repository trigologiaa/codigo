package main

import "fmt"

func main() {
	a := 10
	mostrarDuplicado(a)
	fmt.Println("main.a:", a)
	fmt.Println("main.&a:", &a)
	duplicar(&a) // Le envio la direccion
	fmt.Println("main.a:", a)
	fmt.Println("main.&a:", &a)
}

// a no existe
// b vale a
// Podemos decir que la funcion es inmutable
func mostrarDuplicado(b int) {
	b *= 2
	fmt.Println("mostrarDuplicado.b:", b)
}

func duplicar(c *int) {
	// Esta funcion es mutable
	*c *= 2
	fmt.Println("duplicar.c:", c)
	fmt.Println("duplicar.*c:", *c)
	fmt.Println("duplicar.&c:", &c)
}