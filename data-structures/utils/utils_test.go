package utils

import (
	"math"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestMinimo(test *testing.T) {
	assert.Equal(test, Minimo(1, 2), 1)
	assert.Equal(test, Minimo(2, 1), 1)
	assert.Equal(test, Minimo(1, 1), 1)
}

func TestMaximo(test *testing.T) {
	assert.Equal(test, Maximo(1, 2), 2)
	assert.Equal(test, Maximo(2, 1), 2)
	assert.Equal(test, Maximo(1, 1), 1)
}

func TestComparar(test *testing.T) {
	assert.Equal(test, Comparar(1, 2), -1)
	assert.Equal(test, Comparar(2, 1), 1)
	assert.Equal(test, Comparar(1, 1), 0)
}

func TestCompararNaN(test *testing.T) {
	assert.Equal(test, Comparar(math.NaN(), math.NaN()), 0)
	assert.Equal(test, Comparar(math.NaN(), 1), -1)
	assert.Equal(test, Comparar(1, math.NaN()), 1)
}