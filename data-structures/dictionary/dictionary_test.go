package dictionary

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNuevoDiccionario(test *testing.T) {
	diccionario := NuevoDiccionario[int, int]()
	assert.NotNil(test, diccionario)
	assert.Equal(test, 0, diccionario.Tamaño())
}

func TestContieneElDiccionario(test *testing.T) {
	diccionario := NuevoDiccionario[string, int]()
	diccionario.Insertar("Leo", 55)
	diccionario.Insertar("Lucas", 38)
	assert.True(test, diccionario.Contiene("Leo"))
	assert.True(test, diccionario.Contiene("Lucas"))
	assert.False(test, diccionario.Contiene("Fede"))
}

func TestInsertarEnDiccionario(test *testing.T) {
	diccionario := NuevoDiccionario[string, int]()
	diccionario.Insertar("Leo", 55)
	diccionario.Insertar("Lucas", 38)
	assert.Equal(test, 2, diccionario.Tamaño())
}

func TestInsertarYReemplazarEnDiccionario(test *testing.T) {
	diccionario := NuevoDiccionario[string, int]()
	diccionario.Insertar("Leo", 55)
	diccionario.Insertar("Leo", 38)
	assert.Equal(test, 1, diccionario.Tamaño())
	valor, _ := diccionario.Obtener("Leo")
	assert.Equal(test, 38, valor)
}

func TestObtenerDeDiccionario(test *testing.T) {
	diccionario := NuevoDiccionario[string, int]()
	diccionario.Insertar("Lucas", 35)
	valor, err := diccionario.Obtener("Lucas")
	assert.Equal(test, 35, valor)
	require.NoError(test, err)
	valor, err = diccionario.Obtener("Fede")
	assert.Equal(test, 0, valor)
	assert.EqualError(test, err, "clave inexistente")
}

func TestRemoverDeDiccionario(test *testing.T) {
	diccionario := NuevoDiccionario[string, int]()
	diccionario.Insertar("Leo", 55)
	diccionario.Insertar("Lucas", 38)
	assert.Equal(test, 2, diccionario.Tamaño())
	diccionario.Remover("Leo")
	assert.Equal(test, 1, diccionario.Tamaño())
	assert.True(test, diccionario.Contiene("Lucas"))
	assert.False(test, diccionario.Contiene("Leo"))
}

func TestRemoverNoExistenteDeDiccionario(test *testing.T) {
	diccionario := NuevoDiccionario[string, int]()
	diccionario.Insertar("Leo", 55)
	diccionario.Insertar("Lucas", 38)
	assert.Equal(test, 2, diccionario.Tamaño())
	diccionario.Remover("Fede")
	assert.Equal(test, 2, diccionario.Tamaño())
}

func TestTamañoDeDiccionario(test *testing.T) {
	diccionario := NuevoDiccionario[string, int]()
	diccionario.Insertar("Leo", 55)
	diccionario.Insertar("Lucas", 38)
	assert.Equal(test, 2, diccionario.Tamaño())
}

func TestValoresDeDiccionario(test *testing.T) {
	dic := NuevoDiccionario[int, int]()
	dic.Insertar(1, 2)
	dic.Insertar(3, 4)
	dic.Insertar(5, 6)
	assert.ElementsMatch(test, []int{6, 4, 2}, dic.Valores())
}

func TestClavesDeDiccionario(test *testing.T) {
	diccionario := NuevoDiccionario[int, int]()
	diccionario.Insertar(1, 2)
	diccionario.Insertar(3, 4)
	diccionario.Insertar(5, 6)
	assert.ElementsMatch(test, []int{1, 5, 3}, diccionario.Claves())
}

func TestStringDeDiccionarioVacio(test *testing.T) {
	diccionario := NuevoDiccionario[int, int]()
	assert.Equal(test, "Diccionario: {}", diccionario.String())
}

func TestStringDeDiccionario(test *testing.T) {
	diccionario := NuevoDiccionario[int, int]()
	diccionario.Insertar(1, 2)
	diccionario.Insertar(3, 4)
	posiblesRepresentaciones := []string{
		"Diccionario: {\n  1: 2\n  3: 4\n}",
		"Diccionario: {\n  3: 4\n  1: 2\n}",
	}
	assert.Contains(test, posiblesRepresentaciones, diccionario.String())
}