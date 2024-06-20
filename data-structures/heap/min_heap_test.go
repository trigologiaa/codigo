package heap

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestCrearMonticuloDeMinimoVacio(test *testing.T) {
	monticulo := NuevoMonticuloMinimo[int]()
	assert.Equal(test, 0, monticulo.Tamaño())
}

func TestRemoveMaximoDeMonticuloDeMinimoVacio(test *testing.T) {
	monticulo := NuevoMonticuloMinimo[int]()
	_, err := monticulo.Remover()
	assert.NotNil(test, err)
}

// Gracias a visualgo.net/en/heap
// por la ayuda para preparar este caso de prueba.
//
// Insertando los siguientes elementos en orden:
// 44, 29, 58, 2, 98, 11, 65, 3, 68, 99
//
// El arbol resultante debería ser:
//
//	[02]
//	├── [03]
//	│   ├── [29]
//	│   │   ├── [44]
//	│   │   └── [68]
//	│   └── [98]
//	│       └── [99]
//	└── [11]
//	    ├── [58]
//	    └── [65]
//
// Como arreglo:
// [02, 03, 11, 29, 98, 58, 65, 44, 68, 99].
func TestCrearInsertarYExtraerDeMonticuloDeMinimo(test *testing.T) {
	secuenciaDeInsercion := []int {44, 29, 58, 2, 98, 11, 65, 3, 68, 99}
	ordenEsperadoDespuesDeInsertar := [][]int {
		{44},
		{29, 44},
		{29, 44, 58},
		{2, 29, 58, 44},
		{2, 29, 58, 44, 98},
		{2, 29, 11, 44, 98, 58},
		{2, 29, 11, 44, 98, 58, 65},
		{2, 3, 11, 29, 98, 58, 65, 44},
		{2, 3, 11, 29, 98, 58, 65, 44, 68},
		{2, 3, 11, 29, 98, 58, 65, 44, 68, 99},
	}
	// Verificaciones iniciales
	monticulo := NuevoMonticuloMinimo[int]()
	assert.Equal(test, 0, monticulo.Tamaño())
	// Verificaciones a medida que vamos insertando
	for indice := 0; indice < len(secuenciaDeInsercion); indice++ {
		monticulo.Insertar(secuenciaDeInsercion[indice])
		assert.Equal(test, ordenEsperadoDespuesDeInsertar[indice], monticulo.elementos)
	}
	ordenEsperadoDespuesDeEliminar := [][]int {
		{3, 29, 11, 44, 98, 58, 65, 99, 68},
		{11, 29, 58, 44, 98, 68, 65, 99},
		{29, 44, 58, 99, 98, 68, 65},
		{44, 65, 58, 99, 98, 68},
		{58, 65, 68, 99, 98},
		{65, 98, 68, 99},
		{68, 98, 99},
		{98, 99},
		{99},
		{},
	}
	for indice := 0; indice < len(secuenciaDeInsercion); indice++ {
		_, err := monticulo.Remover()
		assert.Equal(test, ordenEsperadoDespuesDeEliminar[indice], monticulo.elementos)
		assert.NoError(test, err)
	}
}