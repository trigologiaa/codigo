package cambio_recursivo

import "math"

func Cambio(monto int, coins []int) int {
	if monto == 0 {
		return 0
	}
	menor := int(math.MaxInt32)
	for _, coin := range coins {
		if monto >= coin {
			menor = min(menor, Cambio(monto-coin, coins))
		}
	}
	return 1 + menor
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}