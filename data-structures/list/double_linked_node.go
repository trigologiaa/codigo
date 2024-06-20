package list

// DobleNodoEnlazado representa un nodo de una lista enlazada doble.
type DobleNodoEnlazado[T comparable] struct {
	dato 		T
	siguiente	*DobleNodoEnlazado[T]
	anterior 	*DobleNodoEnlazado[T]
}

// NuevoDobleNodoEnlazado crea un nuevo nodo de lista enlazada doble con el dato especificado.
func NuevoDobleNodoEnlazado[T comparable](dato T) *DobleNodoEnlazado[T] {
	return &DobleNodoEnlazado[T]{dato: dato}
}

// EstablecerDato establece el dato almacenado en el nodo.
func (nodo *DobleNodoEnlazado[T]) EstablecerDato(dato T) {
	nodo.dato = dato
}

// Dato devuelve el dato almacenado en el nodo.
func (nodo *DobleNodoEnlazado[T]) Dato() T {
	return nodo.dato
}

// EstablecerSiguiente establece el nodo siguiente al nodo actual.
func (nodo *DobleNodoEnlazado[T]) EstablecerSiguiente(siguiente *DobleNodoEnlazado[T]) {
	nodo.siguiente = siguiente
}

// Siguiente devuelve el nodo siguiente al nodo actual.
func (nodo *DobleNodoEnlazado[T]) Siguiente() *DobleNodoEnlazado[T] {
	return nodo.siguiente
}

// TieneSiguiente devuelve true si el nodo actual tiene asignado un nodo siguiente.
func (nodo *DobleNodoEnlazado[T]) TieneSiguiente() bool {
	return nodo.siguiente != nil
}

// EstablecerAnterior establece el nodo anterior al nodo actual.
func (nodo *DobleNodoEnlazado[T]) EstablecerAnterior(nuevoAnterior *DobleNodoEnlazado[T]) {
	nodo.anterior = nuevoAnterior
}

// Anterior devuelve el nodo anterior al nodo actual.
func (nodo *DobleNodoEnlazado[T]) Anterior() *DobleNodoEnlazado[T] {
	return nodo.anterior
}

// TieneAnterior devuelve true si el nodo actual tiene asignado un nodo previo.
func (nodo *DobleNodoEnlazado[T]) TieneAnterior() bool {
	return nodo.anterior != nil
}