package avltree

import (
	"errors"
	"github.com/trigologiaa/codigo/data-structures/stack"
	"github.com/trigologiaa/codigo/data-structures/types"
)

type AVLInOrderIterador[T types.Ordenado] struct {
	stack *stack.Pila[AVLNodo[T]] // pila de nodos
}

// NewAVLInOrderIterator crea un iterador en orden para un árbol AVL.
//
// Uso:
//
//	avl := avltree.NewAVLTree[int]()
//	iterator := avltree.NewAVLInOrderIterator(avl.GetRoot())
//
// Parámetros:
//   - `root` la raíz del árbol AVL.
//
// Retorna:
//   - un iterador en orden para el árbol AVL.
func NewAVLInOrderIterator[T types.Ordenado](root *AVLNodo[T]) *AVLInOrderIterador[T] {
	iterator := &AVLInOrderIterador[T]{
		stack: stack.NuevaPila[AVLNodo[T]](),
	}
	iterator.stackLeftChildren(root)
	return iterator
}

// stackLeftChildren apila los nodos hijos izquierdos de un nodo.
//
// Parámetros:
//   - `node` el nodo a partir del cual se apilarán los nodos hijos izquierdos.
func (it *AVLInOrderIterador[T]) stackLeftChildren(node *AVLNodo[T]) {
	for node != nil {
		it.stack.Empujar(*node)
		node = node.getLeft()
	}
}

// HasNext indica si hay más elementos para recorrer.
//
// Uso:
//
//	avl := avltree.NewAVLTree[int]()
//	iterator := avltree.NewAVLInOrderIterator(avl.GetRoot())
//	iterator.HasNext()
//
// Retorna:
//   - `true` si hay más elementos para recorrer, `false` en caso contrario.
func (it *AVLInOrderIterador[T]) HasNext() bool {
	return !it.stack.EstaVacia()
}

// Next devuelve el siguiente elemento del recorrido.
//
// Uso:
//
//	avl := avltree.NewAVLTree[int]()
//	iterator := avltree.NewAVLInOrderIterator(avl.GetRoot())
//	iterator.Next()
//
// Retorna:
//   - el siguiente elemento del recorrido.
func (it *AVLInOrderIterador[T]) Next() (T, error) {
	var data T
	if it.stack.EstaVacia() {
		return data, errors.New("no hay más elementos")
	}
	next, _ := it.stack.Tirar()
	if next.getRight() != nil {
		it.stackLeftChildren(next.getRight())
	}
	return next.data, nil
}