package division_y_conquista

import "fmt"

func Potencia(n int, x int) int {
	fmt.Println("n", n, "x", x)
	if x == 0 {
		return 1
	}
	if x == 1 {
		return n
	}
	if x%2 == 0 {
		mediaPotencia := Potencia(n, x/2)
		return mediaPotencia * mediaPotencia
	}
	mediaPotencia := Potencia(n, (x-1)/2)
	return n * mediaPotencia * mediaPotencia
}