package main

import (
	"bufio"
	"fmt"
	"os"
	"github.com/trigologiaa/codigo/data-structures/stack"
)

func InvertirCadena(cadena string) string {
	s := *stack.New[string]()
	salida := ""
	for _, c := range cadena {
		s.Push(string(c))
	}
	for !s.IsEmpty() {
		c, _ := s.Pop()
		salida += c
	}
	return salida
}

func main() {
	fmt.Print("Ingrese una cadena: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	cadena := scanner.Text()
	fmt.Println(InvertirCadena(cadena))
}