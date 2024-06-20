package list

import "fmt"

// ListaCircular implementa una lista enlazada circular genérica.
type ListaCircular[T comparable] struct {
	cabeza	*DobleNodoEnlazado[T]
	tamaño 	int
}

// NuevaListaCircular crea una nueva lista enlazada circular.
//
// Uso:
//
//	list := NuevaListaCircular[int]() // Crea una nueva lista enlazada circular.
func NuevaListaCircular[T comparable]() *ListaCircular[T] {
	return &ListaCircular[T]{}
}

// Cabeza devuelve el primer nodo de la lista.
//
// Uso:
//
//	cabeza := list.Cabeza() // Obtiene el primer nodo de la lista.
//
// Retorna:
//   - el primer nodo de la lista.
func (lista *ListaCircular[T]) Cabeza() *DobleNodoEnlazado[T] {
	return lista.cabeza
}

// Cola devuelve el último nodo de la lista.
//
// Uso:
//
//	tail := list.Cola() // Obtiene el último nodo de la lista.
//
// Retorna:
//   - el último nodo de la lista.
func (lista *ListaCircular[T]) Cola() *DobleNodoEnlazado[T] {
	if lista.tamaño == 0 {
		return nil
	}
	return lista.cabeza.Anterior()
}

// Tamaño devuelve el tamaño de la lista.
//
// Uso:
//
//	tamaño := list.Tamaño() // Obtiene el tamaño de la lista.
//
// Retorna:
//   - el tamaño de la lista.
func (lista *ListaCircular[T]) Tamaño() int {
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
func (lista *ListaCircular[T]) EstaVacia() bool {
	return lista.tamaño == 0
}

// Limpiar elimina todos los elementos de la lista.
//
// Uso:
//
//	list.Limpiar() // Elimina todos los elementos de la lista.
func (lista *ListaCircular[T]) Limpiar() {
	lista.cabeza = nil
	lista.tamaño = 0
}

// Anteponer agrega un nuevo nodo al principio de la lista.
//
// Uso:
//
//	list.Anteponer(10) // Agrega el valor 10 al principio de la lista.
//
// Parámetros:
//   - `dato`: el valor a agregar al principio de la lista.
func (lista *ListaCircular[T]) Anteponer(dato T) {
	nodo := NuevoDobleNodoEnlazado(dato)
	if lista.tamaño == 0 {
		lista.cabeza = nodo
		lista.cabeza.EstablecerSiguiente(lista.cabeza)
		lista.cabeza.EstablecerAnterior(lista.cabeza)
	} else {
		nodo.EstablecerSiguiente(lista.cabeza)
		nodo.EstablecerAnterior(lista.cabeza.Anterior())
		lista.cabeza.Anterior().EstablecerSiguiente(nodo)
		lista.cabeza.EstablecerAnterior(nodo)
		lista.cabeza = nodo
	}
	lista.tamaño++
}

// Adjuntar agrega un nuevo nodo al final de la lista.
//
// Uso:
//
//	list.Adjuntar(10) // Agrega el valor 10 al final de la lista.
//
// Parámetros:
//   - `dato`: el valor a agregar al final de la lista.
func (lista *ListaCircular[T]) Adjuntar(dato T) {
	nodo := NuevoDobleNodoEnlazado(dato)
	if lista.tamaño == 0 {
		lista.cabeza = nodo
		lista.cabeza.EstablecerSiguiente(lista.cabeza)
		lista.cabeza.EstablecerAnterior(lista.cabeza)
	} else {
		nodo.EstablecerSiguiente(lista.cabeza)
		nodo.EstablecerAnterior(lista.cabeza.Anterior())
		lista.cabeza.Anterior().EstablecerSiguiente(nodo)
		lista.cabeza.EstablecerAnterior(nodo)
	}
	lista.tamaño++
}

// Encontrar busca un nodo en la lista.
//
// Uso:
//
//	nodo := list.Encontrar(10) // Busca el valor 10 en la lista.
//
// Parámetros:
//   - `dato`: el valor a buscar en la lista.
//
// Retorna:
//   - el nodo que contiene el valor buscado; `nil` si el valor no se encuentra en la lista.
func (lista *ListaCircular[T]) Encontrar(dato T) *DobleNodoEnlazado[T] {
	if lista.tamaño == 0 {
		return nil
	}
	nodo := lista.cabeza
	for indice := 0; indice < lista.tamaño; indice++ {
		if nodo.Dato() == dato {
			return nodo
		}
		nodo = nodo.Siguiente()
	}
	return nil
}

// RemoverPrimero elimina el primer nodo de la lista.
//
// Uso:
//
//	list.RemoverPrimero() // Elimina el primer nodo de la lista.
func (lista *ListaCircular[T]) RemoverPrimero() {
	if lista.tamaño == 0 {
		return
	}
	if lista.tamaño == 1 {
		lista.cabeza = nil
		lista.tamaño--
		return
	}
	lista.cabeza.Anterior().EstablecerSiguiente(lista.cabeza.Siguiente())
	lista.cabeza.Siguiente().EstablecerAnterior(lista.cabeza.Anterior())
	lista.cabeza = lista.cabeza.Siguiente()
	lista.tamaño--
}

// RemoverUltimo elimina el último nodo de la lista.
//
// Uso:
//
//	list.RemoverUltimo() // Elimina el último nodo de la lista.
func (lista *ListaCircular[T]) RemoverUltimo() {
	if lista.tamaño == 0 {
		return
	}
	if lista.tamaño == 1 {
		lista.cabeza = nil
		lista.tamaño--
		return
	}
	lista.cabeza.Anterior().Anterior().EstablecerSiguiente(lista.cabeza)
	lista.cabeza.EstablecerAnterior(lista.cabeza.Anterior().Anterior())
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
func (lista *ListaCircular[T]) Remover(dato T) {
	nodo := lista.Encontrar(dato)
	if nodo == nil {
		return
	}
	if nodo == lista.cabeza {
		lista.RemoverPrimero()
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
func (lista *ListaCircular[T]) String() string {
	if lista.EstaVacia() {
		return "ListaCircular: ⇢ [] ⇠"
	}
	resultado := "ListaCircular: ⇢ "
	current := lista.Cabeza()
	for {
		resultado += fmt.Sprintf("[%v]", current.Dato())
		if current == lista.Cola() {
			break
		}
		resultado += " ↔ "
		current = current.Siguiente()
	}
	resultado += " ⇠"
	return resultado
}