package combustible

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestDiapositiva(t *testing.T) {
	assert := assert.New(t)
	distancia := 950
	capacidad := 400
	estaciones := []int{200, 375, 550, 750}
	assert.Equal(2, MinimasRecargas(distancia, capacidad, estaciones))
}

func TestLlegaConLaCargaInicial(t *testing.T) {
	assert := assert.New(t)
	distancia := 950
	capacidad := 950
	estaciones := []int{200, 375, 550, 750}
	assert.Equal(0, MinimasRecargas(distancia, capacidad, estaciones))
}

func TestLlegaParandoEnTodas(t *testing.T) {
	assert := assert.New(t)
	distancia := 950
	capacidad := 200
	estaciones := []int{200, 375, 550, 750}
	assert.Equal(4, MinimasRecargas(distancia, capacidad, estaciones))
}

func TestLlegaParandoSolamenteEnLaUltima(t *testing.T) {
	assert := assert.New(t)
	distancia := 950
	capacidad := 800
	estaciones := []int{200, 375, 550, 750, 950}
	assert.Equal(1, MinimasRecargas(distancia, capacidad, estaciones))
}