// Package heap provee una implementación de un heap/montículo binario.
package heap

import (
	"errors"
	"github.com/trigologiaa/codigo/data-structures/types"
	"github.com/trigologiaa/codigo/data-structures/utils"
)

type Monticulo[T any] struct {
	// contenedor de datos
	elementos []T
	// Función de comparación. Para un heap de mínimo,
	// devuelve -1 si a < b, 0 si a == b, 1 si a > b
	// Para un heap de máximo, devuelve 1 si a < b, 0 si a == b, -1 si a > b
	comparar func(a T, b T) int
}

// NuevoMonticuloMinimo crea un nuevo heap binario de mínimos.
//
// Uso:
//
//	heap := heap.NuevoMonticuloMinimo[int]()
//
// Retorna:
//   - un puntero a un heap binario de mínimos.
func NuevoMonticuloMinimo[T types.Ordenado]() *Monticulo[T] {
	return &Monticulo[T] {
		comparar: utils.Comparar[T],
		elementos: make([]T, 0),
	}
}

// NuevoMonticuloMaximo crea un nuevo heap binario de máximos.
//
// Uso:
//
//	heap := heap.NuevoMonticuloMaximo[int]()
//
// Retorna:
//   - un puntero a un heap binario de máximos.
func NuevoMonticuloMaximo[T types.Ordenado]() *Monticulo[T] {
	comparar := func(a T, b T) int {
		return utils.Comparar[T](b, a)
	}
	return &Monticulo[T] {
		comparar: comparar,
		elementos: make([]T, 0),
	}
}

// NuevoMonticuloGenerico crea un nuevo heap binario con una función de comparación personalizada.
//
// Uso:
//
//	heap := heap.NuevoMonticuloGenerico[int](func(a int, b int) int {
//		if a < b {
//			return -1
//		}
//		if a > b {
//			return 1
//		}
//		return 0
//	})
//
// Parámetros:
//   - `comparar` función de comparación personalizada.
//
// Retorna:
//   - un puntero a un heap binario con una función de comparación personalizada.
func NuevoMonticuloGenerico[T any](comparar func(a T, b T) int) *Monticulo[T] {
	return &Monticulo[T] {
		comparar: comparar,
		elementos: make([]T, 0),
	}
}

// Tamaño retorna la cantidad de elementos en el heap.
//
// Uso:
//
//	size := heap.Tamaño()
//
// Retorna:
//   - la cantidad de elementos en el heap.
func (monticulo *Monticulo[T]) Tamaño() int {
	return len(monticulo.elementos)
}

// Insertar agrega un elemento al heap.
//
// Uso:
//
//	heap := heap.NuevoMonticuloMinimo[int]()
//	heap.Insertar(5)
//
// Parámetros:
//
//	elemento: elemento a agregar al heap.
func (monticulo *Monticulo[T]) Insertar(elemento T) {
	monticulo.elementos = append(monticulo.elementos, elemento)
	monticulo.subirMonticulo(len(monticulo.elementos) - 1)
}

// subirMonticulo reordena el heap hacia arriba.
//
// Parámetros:
//   - `indice` índice del elemento a reordenar.
func (monticulo *Monticulo[T]) subirMonticulo(indice int) {
	for indice > 0 {
		padre := (indice - 1) / 2
		if monticulo.comparar(monticulo.elementos[indice], monticulo.elementos[padre]) > 0 {
			break
		}
		monticulo.elementos[indice], monticulo.elementos[padre] = monticulo.elementos[padre], monticulo.elementos[indice]
		indice = padre
	}
}

// Remover elimina y retorna el elemento en la cima del heap.
//
// Uso:
//
//	heap := heap.NuevoMonticuloMinimo[int]()
//	heap.Insertar(5)
//	elemento, _ := heap.Remover()
//
// Retorna:
//   - el elemento en la cima del heap.
func (monticulo *Monticulo[T]) Remover() (T, error) {
	var elemento T
	if monticulo.Tamaño() == 0 {
		return elemento, errors.New("monticulo vacío")
	}
	elemento = monticulo.elementos[0]
	monticulo.elementos[0] = monticulo.elementos[monticulo.Tamaño() - 1]
	monticulo.elementos = monticulo.elementos[:monticulo.Tamaño() - 1]
	monticulo.bajarMonticulo(0)
	return elemento, nil
}

// bajarMonticulo reordena el heap hacia abajo.
//
// Parámetros:
//   - `indice` índice del elemento a reordenar.
func (monticulo *Monticulo[T]) bajarMonticulo(indice int) {
	for {
		izquierda := (2 * indice) + 1
		derecha := (2 * indice) + 2
		masPequeño := indice
		if izquierda < monticulo.Tamaño() && monticulo.comparar(monticulo.elementos[izquierda], monticulo.elementos[masPequeño]) < 0 {
			masPequeño = izquierda
		}
		if derecha < monticulo.Tamaño() && monticulo.comparar(monticulo.elementos[derecha], monticulo.elementos[masPequeño]) < 0 {
			masPequeño = derecha
		}
		if masPequeño == indice {
			break
		}
		monticulo.elementos[indice], monticulo.elementos[masPequeño] = monticulo.elementos[masPequeño], monticulo.elementos[indice]
		indice = masPequeño
	}
}