package avltree

import (
	"fmt"
	"github.com/trigologiaa/codigo/dato-structures/types"
	"github.com/trigologiaa/codigo/dato-structures/utils"
)

// AVLNodo es un nodo del árbol AVL, además del dato y los hijos, registra la altura.
type AVLNodo[T types.Ordenado] struct {
	dato   		T           // dato
	alto 		int         // altura
	izquierda	*AVLNodo[T] // hijo izquierdo
	derecha  	*AVLNodo[T] // hijo derecho
}

// nuevoAVLNodo crea un nuevo nodo AVL con el dato, y los hijos izquierdo y derecho pasados como parámetros.
//
// Parámetros:
//   - dato: dato del nodo.
//   - izquierda: hijo izquierdo.
//   - derecha: hijo derecho.
//
// Retorna:
//   - un puntero al nodo creado.
func nuevoAVLNodo[T types.Ordenado](dato T, izquierda *AVLNodo[T], derecha *AVLNodo[T]) *AVLNodo[T] {
	return &AVLNodo[T]{izquierda: izquierda, derecha: derecha, dato: dato, alto: 0}
}

// ObtenerDato retorna el dato del nodo.
//
// Uso:
//
//	dato := node.ObtenerDato()
//
// Retorna:
//   - el dato del nodo.
func (nodo *AVLNodo[T]) ObtenerDato() T {
	return nodo.dato
}

// string retorna el dato del nodo en formato string.
//
// Retorna:
//   - el dato del nodo en formato string.
func (nodo *AVLNodo[T]) string() string {
	return fmt.Sprintf("%v", nodo.dato)
}

// obtenerIzquierdo retorna el hijo izquierdo del nodo.
//
// Retorna:
//   - el hijo izquierdo del nodo.
func (nodo *AVLNodo[T]) obtenerIzquierdo() *AVLNodo[T] {
	return nodo.izquierda
}

// obtenerDerecho retorna el hijo derecho del nodo.
//
// Retorna:
//   - el hijo derecho del nodo.
func (nodo *AVLNodo[T]) obtenerDerecho() *AVLNodo[T] {
	return nodo.derecha
}

// obtenerAltura retorna la altura del nodo.
//
// Retorna:
//   - la altura del nodo.
func (nodo *AVLNodo[T]) obtenerAltura() int {
	if nodo == nil {
		return -1
	}
	return nodo.alto
}

// obtenerBalance retorna el balance del nodo.
//
// Retorna:
//   - el balance del nodo.
func (nodo *AVLNodo[T]) obtenerBalance() int {
	if nodo == nil {
		return 0
	}
	return nodo.izquierda.obtenerAltura() - nodo.derecha.obtenerAltura()
}

// actualizarAltura actualiza la altura del nodo.
func (nodo *AVLNodo[T]) actualizarAltura() {
	nodo.alto = 1 + utils.Maximo(nodo.izquierda.obtenerAltura(), nodo.derecha.obtenerAltura())
}

// insertar inserta un nuevo nodo en el árbol AVL.
//
// Parámetros:
//   - valor: valor a insertar.
//
// Retorna:
//   - un puntero al nodo insertado.
func (nodo *AVLNodo[T]) insertar(valor T) *AVLNodo[T] {
	// Si el nodo es nil, lo crea
	if nodo == nil {
		return nuevoAVLNodo[T](valor, nil, nil)
	}
	// Primero inserta el nodo como si fuera un BST común
	switch {
	case valor < nodo.dato:
		nodo.izquierda = nodo.izquierda.insertar(valor)
	case valor > nodo.dato:
		nodo.derecha = nodo.derecha.insertar(valor)
	default: // el elemento ya se encuentra en el árbol
		return nodo
	}
	// Actualiza la altura del nodo, y si es necesario, aplica rotaciones
	nodo.actualizarAltura()
	return nodo.aplicarRotacion()
}

// rotarALaDerecha realiza una rotación simple a la derecha.
func (nodo *AVLNodo[T]) rotarALaDerecha() *AVLNodo[T] {
	hijoIzquierdo := nodo.izquierda   // hijoIzquierdo es el hijo izquierdo de nodo
	hijoDerecho := hijoIzquierdo.derecha // hijoDerecho es el hijo derecho de hijoIzquierdo
	// reasignamos los punteros
	hijoIzquierdo.derecha = nodo
	nodo.izquierda = hijoDerecho
	// Actualizamos las alturas
	nodo.actualizarAltura()
	hijoIzquierdo.actualizarAltura()
	return hijoIzquierdo
}

// rotarALaIzquierda realiza una rotación simple a la izquierda.
func (nodo *AVLNodo[T]) rotarALaIzquierda() *AVLNodo[T] {
	hijoDerecho := nodo.derecha // hijoDerecho es el hijo derecho de nodo
	hijoIzquierdo := hijoDerecho.izquierda // hijoIzquierdo es el hijo izquierdo de hijoDerecho
	// reasignamos los punteros
	hijoDerecho.izquierda = nodo
	nodo.derecha = hijoIzquierdo
	// Actualizamos las alturas
	nodo.actualizarAltura()
	hijoDerecho.actualizarAltura()
	return hijoDerecho
}

// remover elimina un nodo del árbol AVL.
//
// Parámetros:
//   - valor: valor a eliminar.
//
// Retorna:
//   - un puntero al nodo eliminado.
func (nodo *AVLNodo[T]) remover(valor T) *AVLNodo[T] {
	if nodo == nil {
		return nodo
	}
	// Primero elimina el nodo como si fuera un BST común
	switch {
	case valor < nodo.dato:
		nodo.izquierda = nodo.izquierda.remover(valor)
	case valor > nodo.dato:
		nodo.derecha = nodo.derecha.remover(valor)
	default:
		if nodo.izquierda == nil {
			return nodo.derecha
		}
		if nodo.derecha == nil {
			return nodo.izquierda
		}
		temp := nodo.derecha.encontrarMinimo()
		nodo.dato = temp.dato
		nodo.derecha = nodo.derecha.remover(temp.dato)
	}
	// Actualiza la altura del nodo, y si es necesario, aplica rotaciones
	nodo.actualizarAltura()
	return nodo.aplicarRotacion()
}

// aplicarRotacion aplica rotaciones al árbol AVL para balancearlo.
//
// Retorna:
//   - un puntero al nodo raiz del sub-arbol, resultante de aplicar las rotaciones.
func (nodo *AVLNodo[T]) aplicarRotacion() *AVLNodo[T] {
	balance := nodo.obtenerBalance()
	// Si |balance| > 1, el árbol está desbalanceado
	// Debemos aplicar rotaciones para balancearlo
	// Desbalanceado a la izquierda -> rotación simple a derecha
	if balance > 1 {
		// Si además el hijo izquierdo está desbalanceado a la derecha,
		// aplicamos una rotación a la izquierda resultando en un
		// desbalanceo izquierda-derecha
		if nodo.izquierda.obtenerBalance() < 0 {
			nodo.izquierda = nodo.izquierda.rotarALaIzquierda()
		}
		return nodo.rotarALaDerecha()
	}
	// Desbalanceado a la derecha -> rotación simple a izquierda
	if balance < -1 {
		// Si además el hijo derecho está desbalanceado a la izquierda,
		// aplicamos una rotación a la derecha resultando en un
		// desbalanceo derecha-izquierda
		if nodo.derecha.obtenerBalance() > 0 {
			nodo.derecha = nodo.derecha.rotarALaDerecha()
		}
		return nodo.rotarALaIzquierda()
	}
	return nodo
}

// encontrarMinimo retorna el nodo con el valor mínimo del árbol AVL.
//
// Retorna:
//   - un puntero al nodo con el valor mínimo.
func (nodo *AVLNodo[T]) encontrarMinimo() *AVLNodo[T] {
	if nodo.izquierda == nil {
		return nodo
	}
	return nodo.izquierda.encontrarMinimo()
}

// encontrarMaximo retorna el nodo con el valor máximo del árbol AVL.
//
// Retorna:
//   - un puntero al nodo con el valor máximo.
func (nodo *AVLNodo[T]) encontrarMaximo() *AVLNodo[T] {
	if nodo.derecha == nil {
		return nodo
	}
	return nodo.derecha.encontrarMaximo()
}

// buscar busca un valor en el árbol AVL.
//
// Parámetros:
//   - k: valor a buscar.
//
// Retorna:
//   - true si el valor se encuentra en el árbol, false en caso contrario.
func (nodo *AVLNodo[T]) buscar(k T) bool {
	if nodo == nil {
		return false
	}
	if nodo.dato > k {
		return nodo.izquierda.buscar(k)
	}
	if nodo.dato < k {
		return nodo.derecha.buscar(k)
	}
	return true
}

// inOrder devuelve una representación en orden del árbol AVL.
//
// Retorna:
//   - una representación en orden del árbol AVL.
func (nodo *AVLNodo[T]) inOrder() string {
	if nodo == nil {
		return ""
	}
	return nodo.izquierda.inOrder() + " " + nodo.string() + " " + nodo.derecha.inOrder()
}