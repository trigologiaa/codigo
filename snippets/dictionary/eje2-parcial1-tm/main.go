package main

import (
	"fmt"
	"github.com/trigologiaa/codigo/data-structures/dictionary"
	"github.com/trigologiaa/codigo/data-structures/list"
	"github.com/trigologiaa/codigo/data-structures/set"
)

// Escribir una función que reciba como parámetro un diccionario cuyas claves son títulos de libros (strings)
// y cada clave tiene asociada un conjunto de autores (string).
// La función debe devolver un diccionario cuyas claves sean los autores y cuyo valor
// una lista enlazada simple de títulos de libros.
func main() {
	entrada := dictionary.NewDictionary[string, set.Set[string]]()
	autores_1 := set.NewListSet("a1", "a2", "a3")
	autores_2 := set.NewListSet("a2", "a4", "a3")
	autores_3 := set.NewListSet("a5", "a2", "a1")
	autores_4 := set.NewListSet("a5", "a2", "a1")
	entrada.Put("l1", autores_1)
	entrada.Put("l2", autores_2)
	entrada.Put("l3", autores_3)
	entrada.Put("l4", autores_4)
	fmt.Println(entrada.String())
	salida := Autores_Libros(*entrada)
	fmt.Println(salida.String())
}

func Autores_Libros(
	entrada dictionary.Dictionary[string, set.Set[string]],
) *dictionary.Dictionary[string, *list.LinkedList[string]] {
	aux := dictionary.NewDictionary[string, *list.LinkedList[string]]()
	for _, titulo := range entrada.Keys() {
		valores, _ := entrada.Get(titulo)
		for _, autor := range valores.Values() {
			var lista *list.LinkedList[string]
			if aux.Contains(autor) {
				lista, _ = aux.Get(autor)
			} else {
				lista = list.NewLinkedList[string]()
			}
			lista.Append(titulo)
			aux.Put(autor, lista)
		}
	}
	return aux
}