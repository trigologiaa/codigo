package main

import (
	"fmt"
	"github.com/trigologiaa/codigo/data-structures/dictionary"
	"github.com/trigologiaa/codigo/data-structures/set"
	"github.com/trigologiaa/codigo/data-structures/stack"
)

func main() {
	fmt.Println("¿La cadena \"()]{}{[()()]()}\" esta balanceada?", Balanceada("[()]{}{[()()]()}"))
	fmt.Println("¿La cadena \"[(])\" esta balanceada?", Balanceada("[(])"))
}

// Reescribir la función que evalúa si una cadena de paréntesis, corchetes y
// llaves está bien balanceada. Devuelve `true` si está bien balanceada y
// `false` si está mal balanceada. Por ejemplo: `()]{}{[()()]()}` debe
// devolver `true`, mientras que `[(])` debe devolver `false`.
// Resolver usando:
// ** una pila para almacenar los simbolos de apertura
// ** un conjunto para almacenar los simbolos de apertura y
// ** un diccionario para recuperar el complemento del simbolo que se dasapila
// Analizar el orden.
func Balanceada(cadena string) bool {
	deAbrir := set.NewListSet("(", "[", "{")
	parcito := dictionary.NewDictionary[string, string]()
	parcito.Put(")", "(")
	parcito.Put("]", "[")
	parcito.Put("}", "{")
	s := stack.New[string]()
	for _, c := range cadena {
		c := string(c)
		if deAbrir.Contains(c) {
			s.Push(c)
		} else { // si no es de abrir es de cerrar
			t, _ := s.Pop()
			p, _ := parcito.Get(c)
			if t != p {
				return false
			}
		}
	}
	return s.IsEmpty()
}