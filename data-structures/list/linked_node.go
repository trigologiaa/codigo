package list

// NodoEnlazado representa un nodo de una lista enlazada simple.
type NodoEnlazado[T comparable] struct {
	dato T
	siguiente *NodoEnlazado[T]
}

// NuevoNodoEnlazado crea un nuevo nodo de lista enlazada con el dato especificado.
//
// Uso:
//
//	node := list.NuevoNodoEnlazado(10) // Crea un nuevo nodo con el dato 10.
//
// Parámetros:
//   - `dato`: el dato a almacenar en el nodo.
func NuevoNodoEnlazado[T comparable](dato T) *NodoEnlazado[T] {
	return &NodoEnlazado[T]{dato: dato}
}

// EstablecerDato establece el dato almacenado en el nodo.
//
// Uso:
//
//	node.EstablecerDato(20) // Establece el dato del nodo a 20.
//
// Parámetros:
//   - `dato`: el dato a almacenar en el nodo.
func (nodo *NodoEnlazado[T]) EstablecerDato(dato T) {
	nodo.dato = dato
}

// Dato devuelve el dato almacenado en el nodo.
//
// Uso:
//
//	dato := node.Dato() // Obtiene el dato almacenado en el nodo.
//
// Retorna:
//   - el dato almacenado en el nodo.
func (nodo *NodoEnlazado[T]) Dato() T {
	return nodo.dato
}

// EstablecerSiguiente establece el nodo siguiente al nodo actual.
//
// Uso:
//
//	node.EstablecerSiguiente(newNode) // Establece el nodo siguiente al nodo actual.
//
// Parámetros:
//   - `nuevoSiguiente`: el nodo siguiente al nodo actual.
func (nodo *NodoEnlazado[T]) EstablecerSiguiente(nuevoSiguiente *NodoEnlazado[T]) {
	nodo.siguiente = nuevoSiguiente
}

// Siguiente devuelve el nodo siguiente al nodo actual.
//
// Uso:
//
//	nextNode := node.Siguiente() // Obtiene el nodo siguiente al nodo actual.
func (nodo *NodoEnlazado[T]) Siguiente() *NodoEnlazado[T] {
	return nodo.siguiente
}

// TieneSiguiente evalúa si el nodo actual tiene asignado un nodo siguiente.
//
// Uso:
//
//	if node.TieneSiguiente() {
//		fmt.Println("El nodo tiene un nodo siguiente.")
//	}
//
// Retorna:
//   - `true` si el nodo tiene un nodo siguiente; `false` en caso contrario.
func (nodo *NodoEnlazado[T]) TieneSiguiente() bool {
	return nodo.siguiente != nil
}