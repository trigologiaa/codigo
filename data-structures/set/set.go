package set

type Conjunto[T comparable] interface {
	Contiene(element T) bool
	Agregar(elements ...T)
	Remover(element T)
	Tamaño() int
	Valores() []T
}