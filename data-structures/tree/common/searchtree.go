// Package common provee la interfaz SearchTree para los árboles de búsqueda.
package common

import "github.com/trigologiaa/codigo/data-structures/types"

type SearchTree[T types.Ordenado] interface {
	Insert(k T)
	Remove(k T)
	IsEmpty() bool
	FindMin() (T, error)
	FindMax() (T, error)
	Search(k T) bool
	Clear()
	String() string
	Iterador() types.Iterador[T]
}