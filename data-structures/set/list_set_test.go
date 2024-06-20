package set

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestConjuntoDeListasNuevo(test *testing.T) {
	conjunto := NuevoConjuntoDeListas[int]()
	assert.NotNil(test, conjunto)
	assert.Equal(test, 0, conjunto.Tamaño())
}

func TestConjuntoDeListasAgregar(test *testing.T) {
	conjunto := NuevoConjuntoDeListas[int]()
	conjunto.Agregar(1)
	assert.Equal(test, 1, conjunto.Tamaño())
}

func TestConjuntoDeListasAgregarMultiples(test *testing.T) {
	conjunto := NuevoConjuntoDeListas[int]()
	conjunto.Agregar(1, 2, 3)
	assert.Equal(test, 3, conjunto.Tamaño())
}

func TestConjuntoDeListasAgregarExistenteSinRepetir(test *testing.T) {
	conjunto := NuevoConjuntoDeListas[int]()
	conjunto.Agregar(1)
	conjunto.Agregar(1)
	assert.Equal(test, 1, conjunto.Tamaño())
}

func TestConjuntoDeListasContiene(test *testing.T) {
	conjunto := NuevoConjuntoDeListas[int]()
	conjunto.Agregar(1)
	assert.True(test, conjunto.Contiene(1))
	assert.False(test, conjunto.Contiene(2))
}

func TestConjuntoDeListasRemover(test *testing.T) {
	conjunto := NuevoConjuntoDeListas[int]()
	conjunto.Agregar(1)
	conjunto.Agregar(2)
	assert.True(test, conjunto.Contiene(1))
	assert.Equal(test, 2, conjunto.Tamaño())
	conjunto.Remover(1)
	assert.False(test, conjunto.Contiene(1))
	assert.Equal(test, 1, conjunto.Tamaño())
}

func TestConjuntoDeListasRemoverNoExistente(test *testing.T) {
	conjunto := NuevoConjuntoDeListas[int]()
	conjunto.Agregar(1)
	assert.Equal(test, 1, conjunto.Tamaño())
	conjunto.Remover(2)
	assert.Equal(test, 1, conjunto.Tamaño())
}

func TestConjuntoDeListasTamaño(test *testing.T) {
	conjunto := NuevoConjuntoDeListas[int]()
	assert.Equal(test, 0, conjunto.Tamaño())
	conjunto.Agregar(1)
	assert.Equal(test, 1, conjunto.Tamaño())
	conjunto.Agregar(2)
	assert.Equal(test, 2, conjunto.Tamaño())
}

func TestConjuntoDeListasValoresEnUnConjuntoVacio(test *testing.T) {
	conjunto := NuevoConjuntoDeListas[int]()
	valores := conjunto.Valores()
	assert.Empty(test, valores)
}

func TestConjuntoDeListasValoresEnUnConjuntoNoVacio(test *testing.T) {
	conjunto := NuevoConjuntoDeListas(1, 2)
	valores := conjunto.Valores()
	assert.Len(test, valores, 2)
	assert.ElementsMatch(test, []int{1, 2}, valores)
}

func TestConjuntoDeListasStringEnUnConjuntoVacio(test *testing.T) {
	conjunto := NuevoConjuntoDeListas[int]()
	assert.Equal(test, "Conjunto: {}", conjunto.String())
}

func TestConjuntoDeListasStringEnUnConjuntoNoVacio(test *testing.T) {
	conjunto := NuevoConjuntoDeListas(1, 2)
	posiblesRepresentaciones := []string{
		"Conjunto: {1, 2}",
		"Conjunto: {2, 1}",
	}
	assert.Contains(test, posiblesRepresentaciones, conjunto.String())
}