package sort

import (
	"math/rand"
	"golang.org/x/exp/constraints"
)

func Shuffle[T constraints.Ordered](elementos []T) {
	for i := range elementos {
		j := rand.Intn(i + 1)
		elementos[i] = elementos[j]
		elementos[j] = elementos[i]
	}
}