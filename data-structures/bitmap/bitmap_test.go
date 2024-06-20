package bitmap

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestMapaDeBitsNuevoTodoEnCero(test *testing.T) {
	mapa := NuevoMapaDeBits()
	assert.Equal(test, uint32(0), mapa.ObtenerMapa())
}

func TestMapaDeBitsEncenderUnBit(test *testing.T) {
	mapa := NuevoMapaDeBits()
	mapa.Encender(1)
	isOn, _ := mapa.EstaEncendido(1)
	assert.True(test, isOn)
}

func TestMapaDeBitsEncenderUnBitEnPosicionNoValida(test *testing.T) {
	mapa := NuevoMapaDeBits()
	err := mapa.Encender(32)
	assert.EqualError(test, err, "posición no válida")
}

func TestMapaDeBitsApagarUnBit(test *testing.T) {
	mapa := NuevoMapaDeBits()
	mapa.Encender(1)
	mapa.Apagar(1)
	isOn, _ := mapa.EstaEncendido(1)
	assert.False(test, isOn)
}

func TestMapaDeBitsApagarUnBitEnPosicionNoValida(test *testing.T) {
	mapa := NuevoMapaDeBits()
	err := mapa.Apagar(32)
	assert.EqualError(test, err, "posición no válida")
}

func TestMapaDeBitsEncenderTodosLosBits(test *testing.T) {
	mapa := NuevoMapaDeBits()
	for i := uint8(0); i < 32; i++ {
		mapa.Encender(i)
	}
	var max uint32 = 0xffffffff
	assert.Equal(test, max, mapa.ObtenerMapa())
}

func TestMapaDeBitsApagarTodosLosBits(test *testing.T) {
	mapa := NuevoMapaDeBits()
	for i := uint8(0); i < 32; i++ {
		mapa.Encender(i)
	}
	for i := uint8(0); i < 32; i++ {
		mapa.Apagar(i)
	}
	assert.Equal(test, uint32(0), mapa.ObtenerMapa())
}

func TestMapaDeBitsEstadoDeUnBit(test *testing.T) {
	mapa := NuevoMapaDeBits()
	mapa.Encender(1)
	isOn, _ := mapa.EstaEncendido(1)
	assert.True(test, isOn)
}

func TestMapaDeBitsEstadoDeUnBitEnPosicionNoValida(test *testing.T) {
	mapa := NuevoMapaDeBits()
	_, err := mapa.EstaEncendido(32)
	assert.EqualError(test, err, "posición no válida")
}

func TestMapaDeBitsEncenderVariasVecesUnMismoBitNoLoApaga(test *testing.T) {
	mapa := NuevoMapaDeBits()
	mapa.Encender(1)
	mapa.Encender(1)
	isOn, _ := mapa.EstaEncendido(1)
	assert.True(test, isOn)
}

func TestMapaDeBitsGetMap(test *testing.T) {
	mapa := NuevoMapaDeBits()
	mapa.Encender(1)
	assert.Equal(test, uint32(2), mapa.ObtenerMapa())
}