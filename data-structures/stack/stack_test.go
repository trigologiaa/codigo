package stack

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestPilaNueva(test *testing.T) {
	pila := NuevaPila[int]()
	assert.NotNil(test, pila)
}

func TestPilaEmpujar(test *testing.T) {
	pila := NuevaPila[int]()
	pila.Empujar(1)
	assert.False(test, pila.EstaVacia())
}

func TestPilaTope(test *testing.T) {
	pila := NuevaPila[int]()
	pila.Empujar(1)
	valor, err := pila.Tope()
	assert.Equal(test, 1, valor)
	assert.NoError(test, err)
}

func TestPilaTopeCuandoEstaVacia(test *testing.T) {
	pila := NuevaPila[int]()
	_, err := pila.Tope()
	assert.EqualError(test, err, "pila vacía")
}

func TestPilaTirar(test *testing.T) {
	pila := NuevaPila[int]()
	pila.Empujar(1)
	valor, err := pila.Tirar()
	assert.Equal(test, 1, valor)
	assert.NoError(test, err)
}

func TestPilaTirarCuandoEstaVacia(test *testing.T) {
	pila := NuevaPila[int]()
	_, err := pila.Tirar()
	assert.EqualError(test, err, "pila vacía")
}

func TestPilaEstaVacia(test *testing.T) {
	pila := NuevaPila[int]()
	assert.True(test, pila.EstaVacia())
	pila.Empujar(1)
	assert.False(test, pila.EstaVacia())
}