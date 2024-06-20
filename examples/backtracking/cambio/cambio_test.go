package cambio

import (
	"fmt"
	"testing"
)

func TestCambio(t *testing.T) {
	monedas := []int{1, 2, 9, 10, 20, 50}
	cantidad := 27
	solucion := CambioDeMoneda(cantidad, monedas)
	fmt.Println(solucion)
	aux := 0
	for i := 0; i < len(solucion); i++ {
		aux += solucion[i]
	}
	if aux != cantidad {
		t.Error("La solucion no es correcta")
	}
}