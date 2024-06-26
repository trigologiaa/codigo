package binarytree

import (
	"errors"
	"github.com/trigologiaa/codigo/data-structures/stack"
	"github.com/trigologiaa/codigo/data-structures/types"
)

// BinaryTreeIterator es un iterador para recorrer un BinaryTree.
type binaryTreeIterator[T types.Ordenado] struct {
	internalStack *stack.Pila[*BinaryNode[T]]
}

// pushLeftNodes apila los nodos izquierdos desde un nodo.
//
// Parámetros:
//   - `node` un puntero a un BinaryNode.
func (it *binaryTreeIterator[T]) pushLeftNodes(node *BinaryNode[T]) {
	for node != nil {
		it.internalStack.Empujar(node)
		node = node.left
	}
}

// HasNext indica si hay un siguiente dato.
//
// Uso:
//
//	bst := tree.NewBinaryTree[int]()
//	// ...
//	it := bst.Iterator()
//	for it.HasNext() {
//		fmt.Println(it.Next())
//	}
//
// Retorna:
//   - true si hay un siguiente nodo, false en caso contrario.
func (it *binaryTreeIterator[T]) HasNext() bool {
	return !it.internalStack.EstaVacia()
}

// Next devuelve el siguiente dato, respetando el recorrido InOrder.
//
// Uso:
//
//	bst := tree.NewBinaryTree[int]()
//	// ...
//	it := bst.Iterator()
//	for it.HasNext() {
//		fmt.Println(it.Next())
//	}
//
// Retorna:
//   - el dato del siguiente nodo.
func (it *binaryTreeIterator[T]) Next() (T, error) {
	if it.internalStack.EstaVacia() {
		var emptyValue T
		return emptyValue, errors.New("árbol vacío")
	}
	node, _ := it.internalStack.Tirar()
	if node.right != nil {
		it.pushLeftNodes(node.right)
	}
	return node.data, nil
}