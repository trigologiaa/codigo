package main

import (
	"errors"
	"fmt"
)

// Define una interfaz de Iterator que define los métodos necesarios para recorrer una matriz.
type Iterator interface {
	HasNext() bool
	Next() (int, error)
}

// Define una estructura de matriz que contiene los datos y las posiciones actuales.
type Matrix struct {
	data [][]int
}

// El constructor de matriz crea una nueva matriz y la devuelve.
func NewMatrix(pepito [][]int) *Matrix {
	return &Matrix{data: pepito}
}

type MatrixIterator struct {
	matrix        Matrix
	filaActual    int
	columnaActual int
}

// El método Iterator devuelve un iterador para la matriz.
func (m *Matrix) Iterator() *MatrixIterator {
	return &MatrixIterator{
		matrix:        *m,
		filaActual:    0,
		columnaActual: 0,
	}
}

// Implementa la interfaz Iterator para la matriz.
func (m *MatrixIterator) HasNext() bool {
	return m.filaActual < len(m.matrix.data) &&
		m.columnaActual < len(m.matrix.data[m.filaActual])
}

// Implementa la interfaz Iterator para la matriz.
func (m *MatrixIterator) Next() (int, error) {
	if !m.HasNext() {
		return 0, errors.New("no hay más elementos")
	}
	value := m.matrix.data[m.filaActual][m.columnaActual]
	// Actualiza las posiciones actuales para el siguiente valor.
	m.columnaActual++
	if m.columnaActual >= len(m.matrix.data[m.filaActual]) {
		m.filaActual++
		m.columnaActual = 0
	}
	return value, nil
}

func main() {
	// Crea una matriz de ejemplo.
	data := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	// Crea un objeto de matriz y un iterador.
	matrix := NewMatrix(data)
	iterator := matrix.Iterator()
	// Utiliza el iterador para recorrer la matriz.
	for iterator.HasNext() {
		value, _ := iterator.Next()
		fmt.Println(value)
	}
	fmt.Println("--------------------")
	// Al pedir más elementos de los que hay, el iterador devuelve un error.
	// Por eso hay que utilizar HasNext() para asegurarse de que hay más elementos.
	val, err := iterator.Next()
	fmt.Println(val, err)
	fmt.Println("--------------------")
	it1 := matrix.Iterator()
	it2 := matrix.Iterator()
	for it1.HasNext() && it2.HasNext() {
		value1, _ := it1.Next()
		if value1%2 == 0 {
			value2, _ := it2.Next()
			fmt.Println(value1, value2)
		}
		fmt.Println(value1)
	}
}