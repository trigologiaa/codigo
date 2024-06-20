package list

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestListaCircularNueva(test *testing.T) {
	lista := NuevaListaCircular[int]()
	assert.NotNil(test, lista)
	assert.Equal(test, 0, lista.Tamaño())
	assert.Nil(test, lista.Cabeza())
	assert.Nil(test, lista.Cola())
}

func TestListaCircularAnteponer(test *testing.T) {
	lista := NuevaListaCircular[string]()
	lista.Anteponer("1")
	assert.Equal(test, 1, lista.Tamaño())
	assert.Equal(test, "1", lista.Cabeza().Dato())
	assert.Equal(test, "1", lista.Cola().Dato())
	lista.Anteponer("2")
	assert.Equal(test, 2, lista.Tamaño())
	assert.Equal(test, "2", lista.Cabeza().Dato())
	assert.Equal(test, "1", lista.Cola().Dato())
}

func TestListaCircularAdjuntar(test *testing.T) {
	lista := NuevaListaCircular[string]()
	lista.Adjuntar("1")
	assert.Equal(test, 1, lista.Tamaño())
	assert.Equal(test, "1", lista.Cabeza().Dato())
	assert.Equal(test, "1", lista.Cola().Dato())
	lista.Adjuntar("2")
	assert.Equal(test, 2, lista.Tamaño())
	assert.Equal(test, "1", lista.Cabeza().Dato())
	assert.Equal(test, "2", lista.Cola().Dato())
}

func TestListaCircularLimpiar(test *testing.T) {
	lista := NuevaListaCircular[int]()
	lista.Adjuntar(1)
	lista.Adjuntar(2)
	lista.Limpiar()
	assert.Equal(test, 0, lista.Tamaño())
	assert.Nil(test, lista.Cabeza())
	assert.Nil(test, lista.Cola())
}

func TestListaCircularEstaVacia(test *testing.T) {
	lista := NuevaListaCircular[int]()
	assert.True(test, lista.EstaVacia())
	lista.Adjuntar(1)
	assert.False(test, lista.EstaVacia())
}

func TestListaCircularTamaño(test *testing.T) {
	lista := NuevaListaCircular[int]()
	assert.Equal(test, 0, lista.Tamaño())
	lista.Adjuntar(1)
	assert.Equal(test, 1, lista.Tamaño())
	lista.Adjuntar(2)
	assert.Equal(test, 2, lista.Tamaño())
}

func TestListaCircularCabeza(test *testing.T) {
	lista := NuevaListaCircular[int]()
	assert.Nil(test, lista.Cabeza())
	lista.Adjuntar(1)
	assert.Equal(test, 1, lista.Cabeza().Dato())
	lista.Adjuntar(2)
	assert.Equal(test, 1, lista.Cabeza().Dato())
}

func TestListaCircularCola(test *testing.T) {
	lista := NuevaListaCircular[int]()
	assert.Nil(test, lista.Cola())
	lista.Adjuntar(1)
	assert.Equal(test, 1, lista.Cola().Dato())
	lista.Adjuntar(2)
	assert.Equal(test, 2, lista.Cola().Dato())
}

func TestListaCircularRemover(test *testing.T) {
	lista := NuevaListaCircular[int]()
	lista.Adjuntar(1)
	lista.Adjuntar(2)
	lista.Adjuntar(3)
	lista.Remover(2)
	assert.Equal(test, 2, lista.Tamaño())
	assert.Equal(test, 1, lista.Cabeza().Dato())
	assert.Equal(test, 3, lista.Cola().Dato())
	lista.Remover(1)
	assert.Equal(test, 1, lista.Tamaño())
	assert.Equal(test, 3, lista.Cabeza().Dato())
	assert.Equal(test, 3, lista.Cola().Dato())
	lista.Remover(3)
	assert.Equal(test, 0, lista.Tamaño())
	assert.Nil(test, lista.Cabeza())
	assert.Nil(test, lista.Cola())
}

func TestListaCircularRemoverUltimo(test *testing.T) {
	lista := NuevaListaCircular[int]()
	lista.Adjuntar(1)
	lista.Adjuntar(2)
	lista.Adjuntar(3)
	lista.RemoverUltimo()
	assert.Equal(test, 2, lista.Tamaño())
	assert.Equal(test, 1, lista.Cabeza().Dato())
	assert.Equal(test, 2, lista.Cola().Dato())
	lista.RemoverUltimo()
	assert.Equal(test, 1, lista.Tamaño())
	assert.Equal(test, 1, lista.Cabeza().Dato())
	assert.Equal(test, 1, lista.Cola().Dato())
	lista.RemoverUltimo()
	assert.Equal(test, 0, lista.Tamaño())
	assert.Nil(test, lista.Cabeza())
	assert.Nil(test, lista.Cola())
}

func TestListaCircularRemoverPrimero(test *testing.T) {
	lista := NuevaListaCircular[int]()
	lista.Adjuntar(1)
	lista.Adjuntar(2)
	lista.Adjuntar(3)
	lista.RemoverPrimero()
	assert.Equal(test, 2, lista.Tamaño())
	assert.Equal(test, 2, lista.Cabeza().Dato())
	assert.Equal(test, 3, lista.Cola().Dato())
	lista.RemoverPrimero()
	assert.Equal(test, 1, lista.Tamaño())
	assert.Equal(test, 3, lista.Cabeza().Dato())
	assert.Equal(test, 3, lista.Cola().Dato())
	lista.RemoverPrimero()
	assert.Equal(test, 0, lista.Tamaño())
	assert.Nil(test, lista.Cabeza())
	assert.Nil(test, lista.Cola())
}

func TestListaCircularRemoverTodoEnListaVacia(test *testing.T) {
	lista := NuevaListaCircular[int]()
	lista.Remover(1)
	assert.Equal(test, 0, lista.Tamaño())
	assert.Nil(test, lista.Cabeza())
	assert.Nil(test, lista.Cola())
	lista.RemoverPrimero()
	assert.Equal(test, 0, lista.Tamaño())
	assert.Nil(test, lista.Cabeza())
	assert.Nil(test, lista.Cola())
	lista.RemoverUltimo()
	assert.Equal(test, 0, lista.Tamaño())
	assert.Nil(test, lista.Cabeza())
	assert.Nil(test, lista.Cola())
}

func TestListaCircularEncontrar(test *testing.T) {
	lista := NuevaListaCircular[int]()
	lista.Adjuntar(1)
	lista.Adjuntar(2)
	lista.Adjuntar(3)
	nodo := lista.Encontrar(2)
	assert.NotNil(test, nodo)
	assert.Equal(test, 2, nodo.Dato())
	nodo = lista.Encontrar(4)
	assert.Nil(test, nodo)
}

func TestListaCircularEncontrarEnListaVacia(test *testing.T) {
	lista := NuevaListaCircular[int]()
	nodo := lista.Encontrar(1)
	assert.Nil(test, nodo)
}

func TestListaCircularStringEnVacio(test *testing.T) {
	lista := NuevaListaCircular[int]()
	assert.Equal(test, "ListaCircular: ⇢ [] ⇠", lista.String())
}

func TestListaCircularString(test *testing.T) {
	lista := NuevaListaCircular[int]()
	lista.Adjuntar(1)
	lista.Adjuntar(2)
	lista.Adjuntar(3)
	assert.Equal(test, "ListaCircular: ⇢ [1] ↔ [2] ↔ [3] ⇠", lista.String())
}