// Package queue implementa una cola genérica sobre un arreglo dinámico.
// Expone la estructura Cola y sus métodos para manipular una cola.
package queue

import "errors"

// Cola implementa una cola genérica sobre un arreglo dinámico.
type Cola[T any] struct {
	dato []T
}

// NuevaCola crea una nueva cola vacía. O(1)
//
// Uso:
//
//	cola := queue.New[int]() // Crea una nueva cola de enteros.
func NuevaCola[T any]() *Cola[T] {
	return &Cola[T]{}
}

// AgregarACola agrega un elemento a la cola. O(1)
//
// Uso:
//
//	cola.AgregarACola(10) // Agrega el entero 10 a la cola.
//
// Parámetros:
//   - `valor`: el elemento a agregar a la cola.
func (cola *Cola[T]) AgregarACola(valor T) {
	cola.dato = append(cola.dato, valor)
}

// SacarDeCola saca un elemento de la cola. O(1)
//
// Uso:
//
//	valor, err := cola.SacarDeCola() // Saca un elemento de la cola.
//
// Retorna:
//   - el elemento sacado de la cola.
//   - un error si la cola está vacía.
func (cola *Cola[T]) SacarDeCola() (T, error) {
	var cabeza T
	if len(cola.dato) == 0 {
		return cabeza, errors.New("cola vacía")
	}
	cabeza = cola.dato[0]
	cola.dato = cola.dato[1:]
	return cabeza, nil
}

// Frente devuelve el elemento del frente de la cola. O(1)
// Esta operación no modifica la cola.
//
// Uso:
//
//	valor, err := cola.Frente() // Obtiene el elemento del frente de la cola.
//
// Retorna:
//   - el elemento del frente de la cola.
//   - un error si la cola está vacía.
func (cola *Cola[T]) Frente() (T, error) {
	var cabeza T
	if len(cola.dato) == 0 {
		return cabeza, errors.New("cola vacía")
	}
	cabeza = cola.dato[0]
	return cabeza, nil
}

// EstaVacia verifica si la cola esta vacia. O(1)
//
// Uso:
//
//	empty := cola.EstaVacia() // Verifica si la cola está vacía.
//
// Retorna:
//   - `true` si la cola está vacía; `false` en caso contrario.
func (cola *Cola[T]) EstaVacia() bool {
	return len(cola.dato) == 0
}