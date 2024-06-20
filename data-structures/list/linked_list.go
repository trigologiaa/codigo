// Package list proporciona implementaciones de listas enlazadas simples, dobles y circulares.
// Expone las estructuras ListaEnlazada, DoubleLinkedList y CircularList y sus métodos para manipular listas.
package list

import "fmt"

// ListaEnlazada se implementa con un nodo que contiene un dato y un puntero al siguiente nodo.
// Los elementos deben ser de un tipo comparable.
type ListaEnlazada[T comparable] struct {
	cabeza	*NodoEnlazado[T]
	cola 	*NodoEnlazado[T]
	tamaño 	int
}

// NuevaListaEnlazada crea una nueva lista vacía.
//
// Uso:
//
//	list := list.NuevaListaEnlazada[int]() // Crea una nueva lista vacía.
func NuevaListaEnlazada[T comparable]() *ListaEnlazada[T] {
	return &ListaEnlazada[T]{}
}

// Cabeza devuelve el primer nodo de la lista.
//
// Uso:
//
//	cabeza := list.Cabeza() // Obtiene el primer nodo de la lista.
//
// Retorna:
//   - el primer nodo de la lista.
func (lista *ListaEnlazada[T]) Cabeza() *NodoEnlazado[T] {
	return lista.cabeza
}

// Cola devuelve el último nodo de la lista.
//
// Uso:
//
//	cola := list.Cola() // Obtiene el último nodo de la lista.
//
// Retorna:
//   - el último nodo de la lista.
func (lista *ListaEnlazada[T]) Cola() *NodoEnlazado[T] {
	return lista.cola
}

// Tamaño devuelve el tamaño de la lista.
//
// Uso:
//
//	tamaño := list.Tamaño() // Obtiene el tamaño de la lista.
//
// Retorna:
//   - el tamaño de la lista.
func (lista *ListaEnlazada[T]) Tamaño() int {
	return lista.tamaño
}

// EstaVacia evalúa si la lista está vacía.
//
// Uso:
//
//	empty := list.EstaVacia() // Verifica si la lista está vacía.
//
// Retorna:
//   - `true` si la lista está vacía; `false` en caso contrario.
func (lista *ListaEnlazada[T]) EstaVacia() bool {
	return lista.tamaño == 0
}

// Limpiar elimina todos los nodos de la lista.
//
// Uso:
//
//	list.Limpiar() // Elimina todos los nodos de la lista.
func (lista *ListaEnlazada[T]) Limpiar() {
	lista.cabeza = nil
	lista.cola = nil
	lista.tamaño = 0
}

// Anteponer inserta un dato al inicio de la lista.
//
// Uso:
//
//	list.Anteponer(10) // Inserta el dato 10 al inicio de la lista.
//
// Parámetros:
//   - `dato`: el dato a insertar en la lista.
func (lista *ListaEnlazada[T]) Anteponer(dato T) {
	nuevoNodo := NuevoNodoEnlazado(dato)
	if lista.EstaVacia() {
		lista.cola = nuevoNodo
	} else {
		nuevoNodo.EstablecerSiguiente(lista.cabeza)
	}
	lista.cabeza = nuevoNodo
	lista.tamaño++
}

// Adjuntar inserta un dato al final de la lista.
//
// Uso:
//
//	list.Adjuntar(10) // Inserta el dato 10 al final de la lista.
//
// Parámetros:
//   - `dato`: el dato a insertar en la lista.
func (lista *ListaEnlazada[T]) Adjuntar(dato T) {
	nuevoNodo := NuevoNodoEnlazado(dato)
	if lista.EstaVacia() {
		lista.cabeza = nuevoNodo
	} else {
		lista.cola.EstablecerSiguiente(nuevoNodo)
	}
	lista.cola = nuevoNodo
	lista.tamaño++
}

// Encontrar busca un dato en la lista, si lo encuentra devuelve el nodo
// correspondiente, si no lo encuentra devuelve nil
//
// Uso:
//
//	nodo := list.Encontrar(10) // Busca el dato 10 en la lista.
//
// Parámetros:
//   - `dato`: el dato a buscar en la lista.
//
// Retorna:
//   - el nodo que contiene el dato; `nil` si el dato no se encuentra.
func (lista *ListaEnlazada[T]) Encontrar(dato T) *NodoEnlazado[T] {
	for actual := lista.cabeza; actual != nil; actual = actual.Siguiente() {
		if actual.Dato() == dato {
			return actual
		}
	}
	return nil
}

// RemoverPrimero elimina el primer nodo de la lista.
//
// Uso:
//
//	list.RemoverPrimero() // Elimina el primer nodo de la lista.
func (lista *ListaEnlazada[T]) RemoverPrimero() {
	if lista.EstaVacia() {
		return
	}
	lista.cabeza = lista.cabeza.Siguiente()
	if lista.cabeza == nil {
		lista.cola = nil
	}
	lista.tamaño--
}

// RemoverUltimo elimina el último nodo de la lista.
//
// Uso:
//
//	list.RemoverUltimo() // Elimina el último nodo de la lista.
func (lista *ListaEnlazada[T]) RemoverUltimo() {
	if lista.EstaVacia() {
		return
	}
	if lista.Tamaño() == 1 {
		lista.cabeza = nil
		lista.cola = nil
	} else {
		actual := lista.cabeza
		for actual.Siguiente() != lista.cola {
			actual = actual.Siguiente()
		}
		actual.EstablecerSiguiente(nil)
		lista.cola = actual
	}
	lista.tamaño--
}

// Remover elimina un la primera aparición de un dato en la lista.
//
// Uso:
//
//	list.Remover(10) // Elimina la primera aparición del dato 10 en la lista.
//
// Parámetros:
//   - `dato`: el dato a eliminar de la lista.
func (lista *ListaEnlazada[T]) Remover(dato T) {
	nodo := lista.Encontrar(dato)
	if nodo == nil {
		return
	}
	if nodo == lista.cabeza {
		lista.RemoverPrimero()
		return
	}
	actual := lista.Cabeza()
	for actual.Siguiente() != nodo {
		actual = actual.Siguiente()
	}
	actual.EstablecerSiguiente(nodo.Siguiente())
	if nodo == lista.cola {
		lista.cola = actual
	}
	lista.tamaño--
}

// String devuelve una representación en cadena de la lista.
//
// Uso:
//
//	fmt.Println(list) // Imprime la representación en cadena de la lista.
//
// Retorna:
//   - una representación en cadena de la lista.
func (lista *ListaEnlazada[T]) String() string {
	if lista.EstaVacia() {
		return "ListaEnlazada: []"
	}
	resultado := "ListaEnlazada: "
	actual := lista.Cabeza()
	for {
		resultado += fmt.Sprintf("[%v]", actual.Dato())
		if !actual.TieneSiguiente() {
			break
		}
		resultado += " → "
		actual = actual.Siguiente()
	}
	return resultado
}