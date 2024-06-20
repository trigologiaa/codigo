package common

import (
	"testing"
	"github.com/trigologiaa/codigo/data-structures/tree/avltree"
	"github.com/trigologiaa/codigo/data-structures/tree/binarytree"
)

// Verifica si BinarySearchTree implementa la interfaz SearchTree.
func TestBSTCompliesTreeInterface(_ *testing.T) {
	var _ SearchTree[int] = binarytree.NewBinarySearchTree[int]()
}

// Verifica si AVLTree implementa la interfaz SearchTree.
func TestAVLTreeCompliesTreeInterface(_ *testing.T) {
	var _ SearchTree[int] = avltree.NewAVLTree[int]()
}