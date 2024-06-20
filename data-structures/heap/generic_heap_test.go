package heap

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

type Persona struct {
	nombre string
	edad   int
}

func personasDeMayorAMenorEdad(axel Persona, bruno Persona) int {
	if axel.edad < bruno.edad {
		return 1
	} else if axel.edad > bruno.edad {
		return -1
	}
	return 0
}

func TestCrearMonticuloVacio(test *testing.T) {
	monticulo := NuevoMonticuloGenerico[Persona](personasDeMayorAMenorEdad)
	assert.Equal(test, 0, monticulo.Tamaño())
}

func TestRemoverDeMonticuloVacio(test *testing.T) {
	monticulo := NuevoMonticuloGenerico[Persona](personasDeMayorAMenorEdad)
	_, err := monticulo.Remover()
	assert.NotNil(test, err)
}

func TestCrearInsertarYExtraerDeMonticulo(test *testing.T) {
	secuenciaDeInsercion := []Persona {
		{"Ana", 44},
		{"Juan", 29},
		{"Pedro", 58},
		{"Maria", 2},
		{"Jose", 98},
		{"Lucia", 11},
		{"Carlos", 65},
		{"Sofia", 3},
		{"Miguel", 68},
		{"Laura", 99},
	}
	ordenEsperadoDespuesDeInsertar := [][]Persona {
		{{"Ana", 44}},
		{{"Ana", 44}, {"Juan", 29}},
		{{"Pedro", 58}, {"Juan", 29}, {"Ana", 44}},
		{{"Pedro", 58}, {"Juan", 29}, {"Ana", 44}, {"Maria", 2}},
		{{"Jose", 98}, {"Pedro", 58}, {"Ana", 44}, {"Maria", 2}, {"Juan", 29}},
		{{"Jose", 98}, {"Pedro", 58}, {"Ana", 44}, {"Maria", 2}, {"Juan", 29}, {"Lucia", 11}},
		{{"Jose", 98}, {"Pedro", 58}, {"Carlos", 65}, {"Maria", 2}, {"Juan", 29}, {"Lucia", 11}, {"Ana", 44}},
		{{"Jose", 98}, {"Pedro", 58}, {"Carlos", 65}, {"Sofia", 3}, {"Juan", 29}, {"Lucia", 11}, {"Ana", 44}, {"Maria", 2}},
		{{"Jose", 98}, {"Miguel", 68}, {"Carlos", 65}, {"Pedro", 58}, {"Juan", 29}, {"Lucia", 11}, {"Ana", 44}, {"Maria", 2}, {"Sofia", 3}},
		{{"Laura", 99}, {"Jose", 98}, {"Carlos", 65}, {"Pedro", 58}, {"Miguel", 68}, {"Lucia", 11}, {"Ana", 44}, {"Maria", 2}, {"Sofia", 3}, {"Juan", 29}},
	}
	// Verificaciones iniciales
	monticulo := NuevoMonticuloGenerico[Persona](personasDeMayorAMenorEdad)
	assert.Equal(test, 0, monticulo.Tamaño())
	// Verificaciones axel medida que vamos insertando
	for indice := 0; indice < len(secuenciaDeInsercion); indice++ {
		monticulo.Insertar(secuenciaDeInsercion[indice])
		assert.Equal(test, ordenEsperadoDespuesDeInsertar[indice], monticulo.elementos)
	}
	ordenEsperadoDespuesDeEliminar := [][]Persona {
		{{"Jose", 98}, {"Miguel", 68}, {"Carlos", 65}, {"Pedro", 58}, {"Juan", 29}, {"Lucia", 11}, {"Ana", 44}, {"Maria", 2}, {"Sofia", 3}},
		{{"Miguel", 68}, {"Pedro", 58}, {"Carlos", 65}, {"Sofia", 3}, {"Juan", 29}, {"Lucia", 11}, {"Ana", 44}, {"Maria", 2}},
		{{"Carlos", 65}, {"Pedro", 58}, {"Ana", 44}, {"Sofia", 3}, {"Juan", 29}, {"Lucia", 11}, {"Maria", 2}},
		{{"Pedro", 58}, {"Juan", 29}, {"Ana", 44}, {"Sofia", 3}, {"Maria", 2}, {"Lucia", 11}},
		{{"Ana", 44}, {"Juan", 29}, {"Lucia", 11}, {"Sofia", 3}, {"Maria", 2}},
		{{"Juan", 29}, {"Sofia", 3}, {"Lucia", 11}, {"Maria", 2}},
		{{"Lucia", 11}, {"Sofia", 3}, {"Maria", 2}},
		{{"Sofia", 3}, {"Maria", 2}},
		{{"Maria", 2}},
		{},
	}
	for indice := 0; indice < len(secuenciaDeInsercion); indice++ {
		_, err := monticulo.Remover()
		assert.Equal(test, ordenEsperadoDespuesDeEliminar[indice], monticulo.elementos)
		assert.NoError(test, err)
	}
}