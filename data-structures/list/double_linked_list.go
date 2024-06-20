package list

import "fmt"

// DobleListaEnlazada implementa una lista enlazada doble genérica.
type DobleListaEnlazada[T comparable] struct {
	cabeza	*DobleNodoEnlazado[T]
	cola 	*DobleNodoEnlazado[T]
	tamaño 	int
}

// NuevaDobleListaEnlazda crea una nueva lista vacía.
//
// Uso:
//
//	list := NuevaDobleListaEnlazda[int]() // Crea una nueva lista enlazada doble.
func NuevaDobleListaEnlazda[T comparable]() *DobleListaEnlazada[T] {
	return &DobleListaEnlazada[T]{}
}

// Cabeza devuelve el primer nodo de la lista.
//
// Uso:
//
//	cabeza := list.Cabeza() // Obtiene el primer nodo de la lista.
//
// Retorna:
//   - el primer nodo de la lista.
func (lista *DobleListaEnlazada[T]) Cabeza() *DobleNodoEnlazado[T] {
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
func (lista *DobleListaEnlazada[T]) Cola() *DobleNodoEnlazado[T] {
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
func (lista *DobleListaEnlazada[T]) Tamaño() int {
	return lista.tamaño
}

// EstaVacia evalúa si la lista está vacía.
//
// Uso:
//
//	empty := list.EstaVacia() // Verifica si la lista está vacía.
//
// Retorna:
//   - true si la lista está vacía; false en caso contrario.
func (lista *DobleListaEnlazada[T]) EstaVacia() bool {
	return lista.tamaño == 0
}

// Limpiar elimina todos los nodos de la lista.
//
// Uso:
//
//	list.Limpiar() // Elimina todos los nodos de la lista.
func (lista *DobleListaEnlazada[T]) Limpiar() {
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
//   - `data`: el dato a insertar al frente de la lista.
func (lista *DobleListaEnlazada[T]) Anteponer(data T) {
	nuevoNodo := NuevoDobleNodoEnlazado[T](data)
	if lista.tamaño == 0 {
		lista.cola = nuevoNodo
	} else {
		nuevoNodo.EstablecerSiguiente(lista.cabeza)
		lista.cabeza.EstablecerAnterior(nuevoNodo)
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
//   - `data`: el dato a insertar al final de la lista.
func (lista *DobleListaEnlazada[T]) Adjuntar(data T) {
	nuevoNodo := NuevoDobleNodoEnlazado[T](data)
	if lista.tamaño == 0 {
		lista.cabeza = nuevoNodo
	} else {
		lista.cola.EstablecerSiguiente(nuevoNodo)
		nuevoNodo.EstablecerAnterior(lista.cola)
	}
	lista.cola = nuevoNodo
	lista.tamaño++
}

// Encontrar busca un dato en la lista
//
// Uso:
//
//	nodo := list.Encontrar(10) // Busca el dato 10 en la lista.
//
// Parámetros:
//   - `data`: el dato a buscar en la lista.
//
// Retorna:
//   - el nodo que contiene el dato buscado; `nil` si no se encuentra.
func (lista *DobleListaEnlazada[T]) Encontrar(data T) *DobleNodoEnlazado[T] {
	for actual := lista.cabeza; actual != nil; actual = actual.Siguiente() {
		if actual.Dato() == data {
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
func (lista *DobleListaEnlazada[T]) RemoverPrimero() {
	if lista.EstaVacia() {
		return
	}
	lista.cabeza = lista.cabeza.Siguiente()
	lista.tamaño--
	if lista.EstaVacia() {
		lista.cola = nil
	} else {
		lista.cabeza.EstablecerAnterior(nil)
	}
}

// RemoverUltimo elimina el último nodo de la lista.
//
// Uso:
//
//	list.RemoverUltimo() // Elimina el último nodo de la lista.
func (lista *DobleListaEnlazada[T]) RemoverUltimo() {
	if lista.EstaVacia() {
		return
	}
	if lista.tamaño == 1 {
		lista.cabeza = nil
		lista.cola = nil
		lista.tamaño = 0
		return
	}
	lista.cola = lista.cola.Anterior()
	lista.cola.EstablecerSiguiente(nil)
	lista.tamaño--
}

// Remover elimina un la primera aparición de un dato en la lista.
//
// Uso:
//
//	list.Remover(10) // Elimina la primera aparición del dato 10 en la lista.
func (lista *DobleListaEnlazada[T]) Remover(data T) {
	nodo := lista.Encontrar(data)
	if nodo == nil {
		return
	}
	if nodo == lista.cabeza {
		lista.RemoverPrimero()
		return
	}
	if nodo == lista.cola {
		lista.RemoverUltimo()
		return
	}
	nodo.Anterior().EstablecerSiguiente(nodo.Siguiente())
	nodo.Siguiente().EstablecerAnterior(nodo.Anterior())
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
func (lista *DobleListaEnlazada[T]) String() string {
	if lista.EstaVacia() {
		return "DobleListaEnlazada: []"
	}
	resultado := "DobleListaEnlazada: "
	actual := lista.Cabeza()
	for {
		resultado += fmt.Sprintf("[%v]", actual.Dato())
		if !actual.TieneSiguiente() {
			break
		}
		resultado += " ↔ "
		actual = actual.Siguiente()
	}
	return resultado
}