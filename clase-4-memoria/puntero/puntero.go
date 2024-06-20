package main

import "fmt"

func main() {
	a := 10
	showDuplicate(a)
	fmt.Println("main.a:", a)
	fmt.Println("main.&a:", &a)
	duplicate(&a) // Le envio la direccion
	fmt.Println("main.a:", a)
	fmt.Println("main.&a:", &a)
}

// a no existe
// b vale a
// Podemos decir que la funcion es inmutable
func showDuplicate(b int) {
	b *= 2
	fmt.Println("showDuplicate.b:", b)
}

func duplicate(c *int) {
	// Esta funcion es mutable
	*c *= 2
	fmt.Println("duplicate.c:", c)
	fmt.Println("duplicate.*c:", *c)
	fmt.Println("duplicate.&c:", &c)
}