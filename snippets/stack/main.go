package main

import (
	"fmt"
	"github.com/trigologiaa/codigo/data-structures/stack"
)

func main() {
	// Crear una nueva pila de enteros
	s := stack.NewStack[int]()
	// Agregar elementos a la pila
	s.Push(1)
	s.Push(2)
	s.Push(3)
	// Verificar si la pila está vacía
	if s.IsEmpty() {
		fmt.Println("La pila está vacía")
	} else {
		fmt.Println("La pila no está vacía")
	}
	// Consultar el elemento en la cima de la pila
	topElement, _ := s.Top()
	fmt.Println("Elemento en la cima de la pila:", topElement)
	// Extraer elementos de la pila
	for !s.IsEmpty() {
		poppedElement, _ := s.Pop()
		fmt.Println("Elemento extraído de la pila:", poppedElement)
	}
	// Verificar si la pila está vacía
	if s.IsEmpty() {
		fmt.Println("La pila está vacía")
	} else {
		fmt.Println("La pila no está vacía")
	}
}