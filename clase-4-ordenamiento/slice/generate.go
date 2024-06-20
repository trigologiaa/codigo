package slice

import (
	"math/rand"
	"time"
)

func GenerarNumeros(tamaño int) []float64 {
	slice := make([]float64, tamaño)
	rand.Seed(time.Now().UnixNano())
	for indice := 0; indice < tamaño; indice++ {
		slice[indice] = float64(rand.Intn(99))
	}
	return slice
}

func GenerarNumerosConsecutivos(tamaño int) []float64 {
	slice := make([]float64, tamaño)
	for indice := 0; indice < tamaño; indice++ {
		slice[indice] = float64(indice + 1)
	}
	return slice
}