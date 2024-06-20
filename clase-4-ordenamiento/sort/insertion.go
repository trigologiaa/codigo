package sort

import (
	"golang.org/x/exp/constraints"
	"time"
)

func Insertion[T constraints.Ordered](elementos []T) {
	for indice1 := 1; indice1 < len(elementos); indice1++ {
		valor := elementos[indice1]
		indice2 := indice1 - 1
		for indice2 >= 0 && elementos[indice2] > valor {
			elementos[indice2 + 1] = elementos[indice2]
			indice2 = indice2 - 1
		}
		elementos[indice2 + 1] = valor
	}
}

func InsertionWithSleep[T constraints.Ordered](elementos []T, sleepMs int, callback func(), color func(...int)) {
	for indice1 := 1; indice1 < len(elementos); indice1++ {
		valor := elementos[indice1]
		indice2 := indice1 - 1
		for indice2 >= 0 && elementos[indice2] > valor {
			color(indice2, indice2 + 1)
			time.Sleep(time.Duration(sleepMs) * time.Millisecond)
			callback()
			elementos[indice2 + 1] = elementos[indice2]
			indice2 = indice2 - 1
			color(indice2 + 1, indice2 + 2)
			time.Sleep(time.Duration(sleepMs) * time.Millisecond)
			callback()
		}
		elementos[indice2 + 1] = valor
		color(indice2 + 1)
		time.Sleep(time.Duration(sleepMs) * time.Millisecond)
		callback()
	}
	color()
}