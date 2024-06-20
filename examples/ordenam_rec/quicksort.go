package ordenamientosRecursivos

// implementa el algoritmo de ordenamiento QuickSort para ordenar un arreglo de enteros de manera ascendente.
func Quicksort(array []int) []int {
	if len(array) < 2 {
		return array // Si el tamaño del arreglo es menor a 2, no hay nada que ordenar, por lo que se retorna
	}
	pivot := array[0]   // Se selecciona el primer elemento como pivote
	i := 1              // Índice para recorrer el arreglo desde el segundo elemento
	j := len(array) - 1 // Índice para recorrer el arreglo desde el último elemento
	// Se realiza el particionamiento del arreglo
	for i < j {
		// Se busca un elemento mayor al pivote desde la izquierda
		for array[i] < pivot && i < len(array)-1 {
			i++
		}
		// Se busca un elemento menor al pivote desde la derecha
		for array[j] > pivot && j > 0 {
			j--
		}
		if i < j {
			// Se intercambian los elementos encontrados
			array[i], array[j] = array[j], array[i]
		}
	}
	// Se coloca el pivote en su posición final
	if array[j] < pivot {
		array[0], array[j] = array[j], array[0]
	}
	// Se realiza la ordenación recursiva de las sublistas
	Quicksort(array[:j])   // Ordenar la sublista izquierda del pivote
	Quicksort(array[j+1:]) // Ordenar la sublista derecha del pivote
	return array
}