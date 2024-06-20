// Package utils provee funciones de utilidad para el resto de los paquetes.
package utils

import "github.com/trigologiaa/codigo/data-structures/types"

// Comparar devuelve
//
//	-1 si x es menor que y,
//	 0 si x es igual a y,
//	+1 si x es mayor que y.
//
// Para tipos de punto flotante, un NaN se considera menor que cualquier valor no-NaN,
// un NaN se considera igual a un NaN, y -0.0 es igual a 0.0.
func Comparar[T types.Ordenado](x, y T) int {
	xNaN := esNaN(x)
	yNaN := esNaN(y)
	if xNaN && yNaN {
		return 0
	}
	if xNaN || x < y {
		return -1
	}
	if yNaN || x > y {
		return +1
	}
	return 0
}

// esNaN informa si x es un NaN sin requerir el paquete math.
// Esto siempre devolverá falso si T no es de punto flotante.
func esNaN[T types.Ordenado](x T) bool {
	return x != x
}

// Minimo devuelve el menor de dos enteros.
//
// Uso:
//
//	a := 5
//	b := 3
//	utils.Minimo(a, b)
//
// Parámetros:
//   - `a` un entero.
//   - `b` un entero.
//
// Retorna:
//   - el menor de los dos enteros.
func Minimo(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Maximo devuelve el mayor de dos enteros.
//
// Uso:
//
//	a := 5
//	b := 3
//	utils.Maximo(a, b)
//
// Parámetros:
//   - `a` un entero.
//   - `b` un entero.
//
// Retorna:
//   - el mayor de los dos enteros.
func Maximo(a, b int) int {
	if a > b {
		return a
	}
	return b
}