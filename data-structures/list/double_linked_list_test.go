package list

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestDobleListaEnlazadaNueva(test *testing.T) {
	lista := NuevaDobleListaEnlazda[int]()
	assert.NotNil(test, lista)
	assert.Equal(test, 0, lista.Tamaño())
	assert.Nil(test, lista.Cabeza())
	assert.Nil(test, lista.Cola())
}

func TestDobleListaEnlazadaAnteponer(test *testing.T) {
	lista := NuevaDobleListaEnlazda[string]()
	lista.Anteponer("1")
	assert.Equal(test, 1, lista.Tamaño())
	assert.Equal(test, "1", lista.Cabeza().Dato())
	assert.Equal(test, "1", lista.Cola().Dato())
	lista.Anteponer("2")
	assert.Equal(test, 2, lista.Tamaño())
	assert.Equal(test, "2", lista.Cabeza().Dato())
	assert.Equal(test, "1", lista.Cola().Dato())
}

func TestDobleListaEnlazadaAdjuntar(test *testing.T) {
	lista := NuevaDobleListaEnlazda[string]()
	lista.Adjuntar("1")
	assert.Equal(test, 1, lista.Tamaño())
	assert.Equal(test, "1", lista.Cabeza().Dato())
	assert.Equal(test, "1", lista.Cola().Dato())
	lista.Adjuntar("2")
	assert.Equal(test, 2, lista.Tamaño())
	assert.Equal(test, "1", lista.Cabeza().Dato())
	assert.Equal(test, "2", lista.Cola().Dato())
}

func TestDobleListaEnlazadaLimpiar(test *testing.T) {
	lista := NuevaDobleListaEnlazda[int]()
	lista.Adjuntar(1)
	lista.Adjuntar(2)
	lista.Limpiar()
	assert.Equal(test, 0, lista.Tamaño())
	assert.Nil(test, lista.Cabeza())
	assert.Nil(test, lista.Cola())
}

func TestDobleListaEnlazadaEstaVacia(test *testing.T) {
	lista := NuevaDobleListaEnlazda[int]()
	assert.True(test, lista.EstaVacia())
	lista.Adjuntar(1)
	assert.False(test, lista.EstaVacia())
}

func TestDobleListaEnlazadaTamaño(test *testing.T) {
	lista := NuevaDobleListaEnlazda[int]()
	assert.Equal(test, 0, lista.Tamaño())
	lista.Adjuntar(1)
	assert.Equal(test, 1, lista.Tamaño())
	lista.Adjuntar(2)
	assert.Equal(test, 2, lista.Tamaño())
}

func TestDobleListaEnlazadaCabeza(test *testing.T) {
	lista := NuevaDobleListaEnlazda[int]()
	assert.Nil(test, lista.Cabeza())
	lista.Adjuntar(1)
	assert.Equal(test, 1, lista.Cabeza().Dato())
	lista.Adjuntar(2)
	assert.Equal(test, 1, lista.Cabeza().Dato())
}

func TestDobleListaEnlazadaCola(test *testing.T) {
	lista := NuevaDobleListaEnlazda[int]()
	assert.Nil(test, lista.Cola())
	lista.Adjuntar(1)
	assert.Equal(test, 1, lista.Cola().Dato())
	lista.Adjuntar(2)
	assert.Equal(test, 2, lista.Cola().Dato())
}

func TestDobleListaEnlazadaRemover(test *testing.T) {
	lista := NuevaDobleListaEnlazda[int]()
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

func TestDobleListaEnlazadaRemoverUltimoElemento(test *testing.T) {
	lista := NuevaDobleListaEnlazda[int]()
	lista.Adjuntar(1)
	lista.Adjuntar(2)
	lista.Remover(2)
	assert.Equal(test, 1, lista.Tamaño())
	assert.Equal(test, 1, lista.Cabeza().Dato())
	assert.Equal(test, 1, lista.Cola().Dato())
}

func TestDobleListaEnlazadaRemoverPrimero(test *testing.T) {
	lista := NuevaDobleListaEnlazda[int]()
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

func TestDobleListaEnlazadaRemoverUltimo(test *testing.T) {
	lista := NuevaDobleListaEnlazda[int]()
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

func TestDobleListaEnlazadaRemoverTodoEnListaVacia(test *testing.T) {
	lista := NuevaDobleListaEnlazda[int]()
	lista.RemoverPrimero()
	assert.Equal(test, 0, lista.Tamaño())
	assert.Nil(test, lista.Cabeza())
	assert.Nil(test, lista.Cola())
	lista.RemoverUltimo()
	assert.Equal(test, 0, lista.Tamaño())
	assert.Nil(test, lista.Cabeza())
	assert.Nil(test, lista.Cola())
	lista.Remover(1)
	assert.Equal(test, 0, lista.Tamaño())
	assert.Nil(test, lista.Cabeza())
	assert.Nil(test, lista.Cola())
}

func TestDobleListaEnlazadaEncontrar(test *testing.T) {
	lista := NuevaDobleListaEnlazda[int]()
	lista.Adjuntar(1)
	lista.Adjuntar(2)
	lista.Adjuntar(3)
	nodo := lista.Encontrar(2)
	assert.NotNil(test, nodo)
	assert.Equal(test, 2, nodo.Dato())
	nodo = lista.Encontrar(4)
	assert.Nil(test, nodo)
}

func TestDobleListaEnlazadaStringEnVacia(test *testing.T) {
	lista := NuevaDobleListaEnlazda[int]()
	assert.Equal(test, "DobleListaEnlazada: []", lista.String())
}

func TestDobleListaEnlazadaString(test *testing.T) {
	lista := NuevaDobleListaEnlazda[int]()
	lista.Adjuntar(1)
	lista.Adjuntar(2)
	lista.Adjuntar(3)
	assert.Equal(test, "DobleListaEnlazada: [1] ↔ [2] ↔ [3]", lista.String())
}