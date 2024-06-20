package cambio

// solucion al problema del cambio de moneda con backtracking
// cantidad: cantidad de dinero a cambiar
// monedas: monedas disponibles para el cambio
// Devuelve un arreglo con la cantidad de monedas de cada tipo


// Función recursiva que realiza el backtracking para encontrar la solución
func backtracking(cantidad int, denominaciones []int, solucion []int, mejorSolucion *[]int) {
	// Si ya encontramos una solución, comparamos si es mejor que la actual
	if cantidad == 0 { // Encontramos una solución
		if len(solucion) < len(*mejorSolucion) || len(*mejorSolucion) == 0 {
			//*mejorSolucion = solucion // No funciona porque se comparte la misma dirección de memoria
			*mejorSolucion = make([]int, len(solucion))
			copy(*mejorSolucion, solucion)
		}
		return
	}
	if len (solucion) >= len(*mejorSolucion) && len(*mejorSolucion) != 0 {
		return
	}
	// Si no encontramos una solución, seguimos explorando todas las posibilidades
	for i := 0; i < len(denominaciones); i++ { //extender
		if denominaciones[i] <= cantidad { // es factible entregar una moneda de este tipo
			nuevaSolucion := append(solucion, denominaciones[i]) //registrar
			backtracking(cantidad-denominaciones[i], denominaciones[i:], nuevaSolucion, mejorSolucion)
		} //borrar implicito porque use nuevaSolucion
	}
}

// Función principal que llama a la función de backtracking
func CambioDeMoneda(cantidad int, denominaciones []int) []int {
	var solucion []int
	var mejorSolucion []int
	backtracking(cantidad, denominaciones, solucion, &mejorSolucion)
	return mejorSolucion
}