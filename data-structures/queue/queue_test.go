package queue

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestColaNueva(test *testing.T) {
	cola := NuevaCola[int]()
	assert.NotNil(test, cola)
	assert.True(test, cola.EstaVacia())
}

func TestColaAgregarACola(test *testing.T) {
	cola := NuevaCola[int]()
	cola.AgregarACola(1)
	assert.False(test, cola.EstaVacia())
}

func TestColaSacarDeCola(test *testing.T) {
	cola := NuevaCola[int]()
	cola.AgregarACola(1)
	cola.AgregarACola(2)
	valor, _ := cola.SacarDeCola()
	assert.Equal(test, 1, valor)
	valor, _ = cola.SacarDeCola()
	assert.Equal(test, 2, valor)
	assert.True(test, cola.EstaVacia())
}

func TestColaSacarDeColaEnColaVacia(test *testing.T) {
	cola := NuevaCola[int]()
	_, err := cola.SacarDeCola()
	assert.EqualError(test, err, "cola vacía")
}

func TestColaFrenteEnColaVacia(test *testing.T) {
	cola := NuevaCola[int]()
	_, err := cola.Frente()
	assert.EqualError(test, err, "cola vacía")
}

func TestColaFrente(test *testing.T) {
	cola := NuevaCola[int]()
	cola.AgregarACola(1)
	cola.AgregarACola(2)
	cola.AgregarACola(3)
	valor, _ := cola.Frente()
	assert.Equal(test, 1, valor)
	cola.SacarDeCola()
	valor, _ = cola.Frente()
	assert.Equal(test, 2, valor)
	cola.SacarDeCola()
	valor, _ = cola.Frente()
	assert.Equal(test, 3, valor)
}