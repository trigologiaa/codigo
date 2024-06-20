package main

import "fmt"

func main() {
	/**
	  con arreglos
	*/
	vector := [5]int{12, 34, 55, 66, 43}
	fmt.Println("Vector: ", vector)
	fmt.Printf("La dirección de memoria del vector es: %p\n", &vector)
	// cada posición del array
	for indice := range vector {
		fmt.Printf("La dirección de memoria del elemento %d del vector es %p\n", indice, &vector[indice])
	}
	/** Otra versión.
	* Analizar la salida. ¿Porqué da diferente que con &vector[indice]?
	Esto pasa porque se crea una copia de cada valor en 'valor', lo cual las direcciones de memoria son copias de la original
	 */
	fmt.Println()
	for indice, valor := range vector {
		fmt.Printf("La dirección de memoria del elemento %d del vector es %d\n", indice, &valor)
	}
	/**
	  con slices
	*/
	particion := vector[2:5]
	fmt.Println("\nPartición: ", particion)
	fmt.Printf("La direccion de memoria de la partición es: %p\n", &particion)
	// cada posición del slice:
	for indice := range particion {
		fmt.Printf("La direccion de memoria del elemento %d de la partición es: %p\n", indice, &particion[indice])
	}
}