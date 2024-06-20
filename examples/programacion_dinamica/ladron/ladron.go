package ladron_recursivo

import (
	"math"
)

func Rob(houses []int) int {
	return rob(houses, 0)
}

func rob(houses []int, current int) int {
	if current >= len(houses) {
		return 0
	}
	return max(houses[current]+rob(houses, current+2), rob(houses, current+1))
}

func max(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}