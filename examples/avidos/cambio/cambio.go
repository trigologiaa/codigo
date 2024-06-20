package cambio

// Cambio de monedas: 1, 2, 5, 10, 20, 50, 100, 200, 500, 1000, 2000, 10000
// Da cambio con el menor número de billetes y monedas posible
// devuelve un map con la cantidad de billetes y monedas de cada denominación

func Cambiar(cantidad int) map[int]int {
	billetes := []int{10000, 2000, 1000, 500, 200, 100, 50, 20, 10, 5, 2, 1}
	cambio := make(map[int]int)
	for _, denominacion := range billetes {
		if cantidad >= denominacion {
			cambio[denominacion] = cantidad / denominacion
			cantidad = cantidad % denominacion
		}
	}
	return cambio
}