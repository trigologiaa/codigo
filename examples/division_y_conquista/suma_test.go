package division_y_conquista

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestSumaElementos(t *testing.T) {
	arreglo := []int{1, 2, 3, 4, 5, 6}
	assert.Equal(t, 21, SumaArray(arreglo))
}

func TestSumaVacio(t *testing.T) {
	arreglo := []int{}
	assert.Equal(t, 0, SumaArray(arreglo))
}

func TestSumaOpuestos(t *testing.T) {
	arreglo := []int{-2, -1, 0, 1, 2}
	assert.Equal(t, 0, SumaArray(arreglo))
}