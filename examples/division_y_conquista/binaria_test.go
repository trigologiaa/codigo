package division_y_conquista

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestBuscandoElementoExistente(t *testing.T) {
	array := []int{1, 2, 3, 4, 5, 8, 9, 12, 27, 50, 70}
	pos := BusquedaBinaria(array, 27)
	assert.Equal(t, 8, pos, "el elemento existe en la posicion 8")
}

func TestBuscandoElementoInexistente(t *testing.T) {
	array := []int{1, 2, 3, 4, 5, 8, 9, 12, 27, 50, 70}
	pos := BusquedaBinaria(array, 23)
	assert.Equal(t, -1, pos, "el elemento no existe en el array")
}