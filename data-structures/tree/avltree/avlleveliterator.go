package avltree

import (
	"errors"
	"github.com/trigologiaa/codigo/data-structures/queue"
	"github.com/trigologiaa/codigo/data-structures/types"
)

type AVLLevelIterator[T types.Ordenado] struct {
	queue *queue.Cola[AVLNodo[T]] // cola de nodos
}

// NewAVLLevelIterator crea un iterador de nivel para un árbol AVL.
//
// Uso:
//
//	avl := avltree.NewAVLTree[int]()
//	iterator := avltree.NewAVLLevelIterator(avl.GetRoot())
//
// Parámetros:
//   - `root` la raíz del árbol AVL.
//
// Retorna:
//   - un iterador de nivel para el árbol AVL.
func NewAVLLevelIterator[T types.Ordenado](root *AVLNodo[T]) *AVLLevelIterator[T] {
	iterator := &AVLLevelIterator[T]{
		queue: queue.NuevaCola[AVLNodo[T]](),
	}
	if root != nil {
		iterator.queue.AgregarACola(*root)
	}
	return iterator
}

// HasNext indica si hay más elementos para recorrer.
//
// Uso:
//
//	avl := avltree.NewAVLTree[int]()
//	iterator := avltree.NewAVLLevelIterator(avl.GetRoot())
//	iterator.HasNext()
//
// Retorna:
//   - `true` si hay más elementos para recorrer, `false` en caso contrario.
func (it *AVLLevelIterator[T]) HasNext() bool {
	return !it.queue.EstaVacia()
}

// Next devuelve el siguiente elemento del recorrido.
//
// Uso:
//
//	avl := avltree.NewAVLTree[int]()
//	iterator := avltree.NewAVLLevelIterator(avl.GetRoot())
//	iterator.Next()
//
// Retorna:
//   - el siguiente elemento del recorrido.
func (it *AVLLevelIterator[T]) Next() (T, error) {
	var data T
	if it.queue.EstaVacia() {
		return data, errors.New("no hay más elementos")
	}
	next, _ := it.queue.SacarDeCola()
	if next.getLeft() != nil {
		it.queue.AgregarACola(*next.getLeft())
	}
	if next.getRight() != nil {
		it.queue.AgregarACola(*next.getRight())
	}
	return next.data, nil
}