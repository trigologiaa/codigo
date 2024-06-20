package reinas

import (
	"fmt"
	"testing"
)

func esSolucion(solucion []int) bool {
	for i := 0; i < len(solucion); i++ {
		for j := i + 1; j < len(solucion); j++ {
			if solucion[i] == solucion[j] || j-i == solucion[j]-solucion[i] || j-i == solucion[i]-solucion[j] {
				return false
			}
		}
	}
	return true
}

func TestNReinas(t *testing.T) {
	n := 12
	solucion := NReinas(n)
	fmt.Println(solucion)
	if len(solucion) != n {
		t.Error("La solucion no es correcta")
	}
	if !esSolucion(solucion) {
		t.Error("La solucion no es correcta")
	}
}