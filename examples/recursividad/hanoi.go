package recursividad

import "fmt"

func Mover(n int, origen string, aux string, destino string) {
	if n == 1 {
		fmt.Println("disco", n, origen, "->", destino)
	} else {
		Mover(n-1, origen, destino, aux)
		fmt.Println("disco", n, origen, "->", destino)
		Mover(n-1, aux, origen, destino)
	}
}