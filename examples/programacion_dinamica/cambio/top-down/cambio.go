package cambio_top_down

import "math"

func Cambio(monto int, coins []int) int {
	return cambio(monto, coins, make(map[int]int))
}

func cambio(monto int, coins []int, mem map[int]int) int {
	if monto == 0 {
		return 0
	}
	if result, ok := mem[monto]; ok {
		return result
	}
	menor := int(math.MaxInt32)
	for _, coin := range coins {
		if monto >= coin {
			menor = min(menor, cambio(monto-coin, coins, mem))
		}
	}
	mem[monto] = 1 + menor
	return mem[monto]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}