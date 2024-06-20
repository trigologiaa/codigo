package reinas

// El problema de N reinas
// consiste en colocar n reinas en un tablero de ajedrez de nxn
// de tal forma que no se amenacen entre ellas

// Función que implementa la solución al problema de las N reinas
// Recibe como parámetro el número de reinas a colocar
// Retorna un slice con la solución
func NReinas(n int) []int {
	var reinas []int
	var solucion []int
	backtracking(n, 0, reinas, &solucion)
	return solucion
}

// Verifica si una reina puede ser colocada en (fila, columna)
// dada una configuración de reinas ya colocadas
func esFactible(fila int, columna int, reinas []int) bool {
	for i := 0; i < len(reinas); i++ {
		if reinas[i] == columna || fila+columna == i+reinas[i] || fila-columna == i-reinas[i] {
			return false
		}
	}
	return true
}

// Función recursiva que realiza el backtracking
// Parámetros: el número de reinas a colocar, la fila actual,
// la configuración de reinas ya colocadas.
// Retorna la solución encontrada
func backtracking(n int, fila int, reinas []int, solucion *[]int) {
	if fila == n { //Colocamos n reinas en el tablero
		//*solucion = reinas 
		*solucion = make([]int, len(reinas))
		copy(*solucion, reinas)
		return
	}
	//Extender
	for columna := 0; columna < n; columna++ {
		if esFactible(fila, columna, reinas) {
			nuevaSolucion := append(reinas, columna) //Registrar
			backtracking(n, fila+1, nuevaSolucion, solucion)
		}
	}
}