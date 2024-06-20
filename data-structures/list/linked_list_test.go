package list

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestListaEnlazadaNueva(test *testing.T) {
	lista := NuevaListaEnlazada[int]()
	assert.NotNil(test, lista)
	assert.Equal(test, 0, lista.Tamaño())
	assert.Nil(test, lista.Cabeza())
	assert.Nil(test, lista.Cola())
}

func TestListaEnlazadaAnteponerEnVacia(test *testing.T) {
	lista := NuevaListaEnlazada[string]()
	lista.Anteponer("1")
	assert.Equal(test, 1, lista.Tamaño())
	assert.Equal(test, "1", lista.Cabeza().Dato())
	assert.Equal(test, "1", lista.Cola().Dato())
}

func TestListaEnlazadaAnteponerEnNoVacia(test *testing.T) {
	lista := NuevaListaEnlazada[string]()
	lista.Adjuntar("1")
	lista.Anteponer("2")
	assert.Equal(test, 2, lista.Tamaño())
	assert.Equal(test, "2", lista.Cabeza().Dato())
	assert.Equal(test, "1", lista.Cola().Dato())
}

func TestListaEnlazadaAdjuntarEnVacia(test *testing.T) {
	lista := NuevaListaEnlazada[string]()
	lista.Adjuntar("1")
	assert.Equal(test, 1, lista.Tamaño())
	assert.Equal(test, "1", lista.Cabeza().Dato())
	assert.Equal(test, "1", lista.Cola().Dato())
}

func TestListaEnlazadaAdjuntarEnNoVacia(test *testing.T) {
	lista := NuevaListaEnlazada[string]()
	lista.Adjuntar("1")
	lista.Adjuntar("2")
	assert.Equal(test, 2, lista.Tamaño())
	assert.Equal(test, "1", lista.Cabeza().Dato())
	assert.Equal(test, "2", lista.Cola().Dato())
}

func TestListaEnlazadaLimpiar(test *testing.T) {
	lista := NuevaListaEnlazada[int]()
	lista.Adjuntar(1)
	lista.Adjuntar(2)
	lista.Limpiar()
	assert.Equal(test, 0, lista.Tamaño())
	assert.Nil(test, lista.Cabeza())
	assert.Nil(test, lista.Cola())
}

func TestListaEnlazadaEstaVacia(test *testing.T) {
	lista := NuevaListaEnlazada[int]()
	assert.True(test, lista.EstaVacia())
	lista.Adjuntar(1)
	assert.False(test, lista.EstaVacia())
}

func TestListaEnlazadaTamaño(test *testing.T) {
	lista := NuevaListaEnlazada[int]()
	assert.Equal(test, 0, lista.Tamaño())
	lista.Adjuntar(1)
	assert.Equal(test, 1, lista.Tamaño())
	lista.Adjuntar(2)
	assert.Equal(test, 2, lista.Tamaño())
}

func TestListaEnlazadaCabeza(test *testing.T) {
	lista := NuevaListaEnlazada[int]()
	assert.Nil(test, lista.Cabeza())
	lista.Adjuntar(1)
	assert.Equal(test, 1, lista.Cabeza().Dato())
	lista.Adjuntar(2)
	assert.Equal(test, 1, lista.Cabeza().Dato())
}

func TestListaEnlazadaCola(test *testing.T) {
	lista := NuevaListaEnlazada[int]()
	assert.Nil(test, lista.Cola())
	lista.Adjuntar(1)
	assert.Equal(test, 1, lista.Cola().Dato())
	lista.Adjuntar(2)
	assert.Equal(test, 2, lista.Cola().Dato())
}

func TestListaEnlazadaEncontrar(test *testing.T) {
	lista := NuevaListaEnlazada[int]()
	lista.Adjuntar(1)
	lista.Adjuntar(2)
	lista.Adjuntar(3)
	nodo := lista.Encontrar(2)
	assert.NotNil(test, nodo)
	assert.Equal(test, 2, nodo.Dato())
	nodo = lista.Encontrar(4)
	assert.Nil(test, nodo)
}

func TestListaEnlazadaRemoverPrimero(test *testing.T) {
	lista := NuevaListaEnlazada[int]()
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

func TestListaEnlazadaRemoverUltimo(test *testing.T) {
	lista := NuevaListaEnlazada[int]()
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

func TestListaEnlazadaRemover(test *testing.T) {
	lista := NuevaListaEnlazada[int]()
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

func TestListaEnlazadaRemoverEnUltimoElemento(test *testing.T) {
	lista := NuevaListaEnlazada[int]()
	lista.Adjuntar(1)
	lista.Adjuntar(2)
	lista.Adjuntar(3)
	lista.Adjuntar(4)
	lista.Remover(4)
	assert.Equal(test, 3, lista.Tamaño())
	assert.Equal(test, 1, lista.Cabeza().Dato())
	assert.Equal(test, 3, lista.Cola().Dato())
}

func TestListaEnlazadaRemoverNoExistente(test *testing.T) {
	lista := NuevaListaEnlazada[int]()
	lista.Adjuntar(1)
	lista.Adjuntar(2)
	lista.Adjuntar(3)
	lista.Remover(4)
	assert.Equal(test, 3, lista.Tamaño())
	assert.Equal(test, 1, lista.Cabeza().Dato())
	assert.Equal(test, 3, lista.Cola().Dato())
}

func TestListaEnlazadaRemoverPrimeroEnVacia(test *testing.T) {
	lista := NuevaListaEnlazada[int]()
	lista.RemoverPrimero()
	assert.Equal(test, 0, lista.Tamaño())
}

func TestListaEnlazadaRemoverUltimoEnVacia(test *testing.T) {
	lista := NuevaListaEnlazada[int]()
	lista.RemoverUltimo()
	assert.Equal(test, 0, lista.Tamaño())
}

func TestListaEnlazadaStringEnVacia(test *testing.T) {
	lista := NuevaListaEnlazada[int]()
	assert.Equal(test, "ListaEnlazada: []", lista.String())
}

func TestListaEnlazadaString(test *testing.T) {
	lista := NuevaListaEnlazada[int]()
	lista.Adjuntar(1)
	lista.Adjuntar(2)
	lista.Adjuntar(3)
	assert.Equal(test, "ListaEnlazada: [1] → [2] → [3]", lista.String())
}