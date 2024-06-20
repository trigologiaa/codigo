// Package stack proporciona una implementación de una pila genérica sobre un arreglo dinámico.
// Expone la estructura y funciones básicas para manipular una pila.
package stack

import "errors"

// Pila proporciona una pila cuyos elementos son de un tipo genérico.
// La implementación se basa en un arreglo dinámico.
type Pila[T any] struct {
	dato []T
}

// NuevaPila crea una nueva pila vacía.
//
// Uso:
//
//	pila := stack.New[int]() // Crea una pila de enteros.
func NuevaPila[T any]() *Pila[T] {
	return &Pila[T]{}
}

// Empujar agrega un elemento a la pila.
//
// Uso:
//
//	pila.Empujar(10) // Agrega el entero 10 a la pila.
func (pila *Pila[T]) Empujar(elemento T) {
	pila.dato = append(pila.dato, elemento)
}

// Tirar remueve y retorna el elemento en el tope de la pila.
// Si la pila está vacía, retorna un error.
//
// Uso:
//
//	if elemento, err := pila.Tirar(); err != nil {
//		fmt.Println(err)
//	} else {
//		fmt.Println(elemento)
//	}
func (pila *Pila[T]) Tirar() (T, error) {
	var elemento T
	if pila.EstaVacia() {
		return elemento, errors.New("pila vacía")
	}
	elemento = pila.dato[len(pila.dato) - 1]
	pila.dato = pila.dato[:len(pila.dato) - 1]
	return elemento, nil
}

// Tope retorna el elemento en el tope de la pila.
// Si la pila está vacía, retorna un error.
//
// Uso:
//
//	if elemento, err := pila.Tope(); err != nil {
//		fmt.Println(err)
//	} else {
//		fmt.Println(elemento)
//	}
func (pila *Pila[T]) Tope() (T, error) {
	var elemento T
	if pila.EstaVacia() {
		return elemento, errors.New("pila vacía")
	}
	elemento = pila.dato[len(pila.dato) - 1]
	return elemento, nil
}

// EstaVacia retorna true si la pila está vacía.
//
// Uso:
//
//	if pila.EstaVacia() {
//		fmt.Println("La pila está vacía")
//	}
func (pila *Pila[T]) EstaVacia() bool {
	return len(pila.dato) == 0
}