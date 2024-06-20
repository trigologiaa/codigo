// Package conjunto proporciona una implementación de un conjunto genérico basado en una lista enlazada simple.
// Expone la estructura ConjuntoDeListas y sus métodos para manipular un conjunto.
package set

import (
	"fmt"
	"github.com/trigologiaa/codigo/data-structures/list"
)

// ConjuntoDeListas implementa un conjunto sobre una lista enlazada simple.
type ConjuntoDeListas[T comparable] struct {
	elementos *list.ListaEnlazada[T]
}

// NuevoConjuntoDeListas crea un nuevo conjunto vacío y lo inicializa con los elementos especificados.
//
// Uso:
//
//	s1 := conjunto.NuevoConjuntoDeListas[int]() // Crea un nuevo conjunto vacío.
//	s2 := conjunto.NuevoConjuntoDeListas[int](1, 2, 3) // Crea un nuevo conjunto con los elementos 1, 2 y 3.
//
// Parámetros:
//   - `elementos`: los elementos con los que inicializar el conjunto.
func NuevoConjuntoDeListas[T comparable](elementos ...T) *ConjuntoDeListas[T] {
	conjunto := &ConjuntoDeListas[T]{list.NuevaListaEnlazada[T]()}
	conjunto.Agregar(elementos...)
	return conjunto
}

// Contiene verifica si el conjunto contiene el elemento especificado.
//
// Uso:
//
//	if conjunto.Contiene(10) {
//		fmt.Println("El conjunto contiene el elemento 10.")
//	}
//
// Parámetros:
//   - `elemento`: el elemento a buscar en el conjunto.
//
// Retorna:
//   - `true` si el conjunto contiene el elemento; `false` en caso contrario.
func (conjunto *ConjuntoDeListas[T]) Contiene(elemento T) bool {
	return conjunto.elementos.Encontrar(elemento) != nil
}

// Agregar agrega los elementos especificados al conjunto.
//
// Uso:
//
//	conjunto.Agregar(10, 20, 30) // Agrega los elementos 10, 20 y 30 al conjunto.
//
// Parámetros:
//   - `elementos`: los elementos a agregar al conjunto.
func (conjunto *ConjuntoDeListas[T]) Agregar(elementos ...T) {
	for _, elemento := range elementos {
		if !conjunto.Contiene(elemento) {
			conjunto.elementos.Adjuntar(elemento)
		}
	}
}

// Remover elimina el elemento especificado del conjunto.
//
// Uso:
//
//	conjunto.Remover(10) // Elimina el elemento 10 del conjunto.
//
// Parámetros:
//   - `elemento`: el elemento a eliminar del conjunto.
func (conjunto *ConjuntoDeListas[T]) Remover(elemento T) {
	conjunto.elementos.Remover(elemento)
}

// Tamaño devuelve la cantidad de elementos en el conjunto.
//
// Uso:
//
//	size := conjunto.Tamaño() // Obtiene la cantidad de elementos en el conjunto.
//
// Retorna:
//   - la cantidad de elementos en el conjunto.
func (conjunto *ConjuntoDeListas[T]) Tamaño() int {
	return conjunto.elementos.Tamaño()
}

// Valores devuelve los elementos del conjunto.
//
// Uso:
//
//	valores := conjunto.Valores() // Obtiene los elementos del conjunto como un slice.
//
// Retorna:
//   - los elementos del conjunto como un slice.
func (conjunto *ConjuntoDeListas[T]) Valores() []T {
	var valores []T
	for nodo := conjunto.elementos.Cabeza(); nodo != nil; nodo = nodo.Siguiente() {
		valores = append(valores, nodo.Dato())
	}
	return valores
}

// String devuelve una representación en cadena del conjunto.
//
// Uso:
//
//	fmt.Println(conjunto) // Muestra el conjunto como una cadena.
//
// Retorna:
//   - una representación en cadena del conjunto.
func (conjunto *ConjuntoDeListas[T]) String() string {
	String := "Conjunto: {"
	for indice, valor := range conjunto.Valores() {
		if indice > 0 {
			String += ", "
		}
		String += fmt.Sprintf("%v", valor)
	}
	String += "}"
	return String
}