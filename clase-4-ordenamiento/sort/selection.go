package sort

import (
	"time"
	"golang.org/x/exp/constraints"
)

func Selection[T constraints.Ordered](elementos []T) {
	ultimo := len(elementos) - 1
	for indice1 := 0; indice1 < ultimo; indice1++ {
		aMinimo := elementos[indice1]
		iMinimo := indice1
		for indice2 := indice1 + 1; indice2 < len(elementos); indice2++ {
			if elementos[indice2] < aMinimo {
				aMinimo = elementos[indice2]
				iMinimo = indice2
			}
		}
		elementos[indice1], elementos[iMinimo] = aMinimo, elementos[indice1]
	}
}

func SelectionWithSleep[T constraints.Ordered](elementos []T, sleepMs int, callback func(), color func(...int)) {
	ultimo := len(elementos) - 1
	for indice1 := 0; indice1 < ultimo; indice1++ {
		aMinimo := elementos[indice1]
		iMinimo := indice1
		for indice2 := indice1 + 1; indice2 < len(elementos); indice2++ {
			color(indice2, iMinimo)
			time.Sleep(time.Duration(sleepMs) * time.Millisecond)
			callback()
			if elementos[indice2] < aMinimo {
				aMinimo = elementos[indice2]
				iMinimo = indice2
			}
			// color(indice2, iMinimo)
			time.Sleep(time.Duration(sleepMs) * time.Millisecond)
			callback()
		}
		elementos[indice1], elementos[iMinimo] = aMinimo, elementos[indice1]
		color(indice1)
		time.Sleep(time.Duration(sleepMs) * time.Millisecond)
		callback()
	}
	color()
}