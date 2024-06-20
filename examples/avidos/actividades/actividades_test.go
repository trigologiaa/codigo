package actividades

import (
	"fmt"
	"testing"
)

func TestActividades(t *testing.T) {
	actividades := []Actividad{
		{"a1", 1, 4},
		{"a2", 3, 5},
		{"a3", 0, 6},
		{"a4", 5, 7},
		{"a5", 3, 9},
		{"a6", 5, 9},
		{"a7", 6, 10},
		{"a8", 8, 11},
		{"a9", 8, 12},
		{"a10", 2, 14},
		{"a11", 12, 15},
	}
	fmt.Println("Actividades: ", actividades)
	seleccionadas := SelectorActividadesIterativo(actividades)
	fmt.Println("Actividades seleccionadas: ", seleccionadas)
	if len(seleccionadas) != 4 {
		t.Error("Se esperaban 4 actividades seleccionadas")
	}
}