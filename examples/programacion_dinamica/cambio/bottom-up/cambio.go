package cambio_bottom_up

import "math"

func Cambio(monto int, coins []int) int {
	n := monto
	table := make([]int, n+1)
	table[0] = 0
	for x := 1; x <= n; x++ {
		menor := int(math.MaxInt32)
		for _, coin := range coins {
			if x >= coin {
				menor = min(menor, table[x-coin])
			}
		}
		table[x] = 1 + menor
	}
	// fmt.Println(table)
	return table[n]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}