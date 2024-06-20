package division_y_conquista

func BusquedaBinaria(array []int, x int) int {
	return busquedaBinaria(array, 0, len(array)-1, x)
}

func busquedaBinaria(array []int, inicio int, fin int, x int) int {
	if inicio > fin {
		return -1
	}
	medio := (inicio + fin) / 2
	if array[medio] > x {
		return busquedaBinaria(array, inicio, medio-1, x)
	}
	if array[medio] < x {
		return busquedaBinaria(array, medio+1, fin, x)
	}
	return medio
}