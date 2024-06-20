package main

import "fmt"

func main() {
	/**
	  con arreglos
	*/
	vectorcito := [5]int{12, 34, 55, 66, 43}
	fmt.Println("vectorcito: ", vectorcito)
	fmt.Printf("La dirección de memoria de vectorcito es: %p\n", &vectorcito)
	// cada posición del array
	for i := range vectorcito {
		fmt.Printf(
			"La dirección de memoria del elemento %d de vectorcito es %p\n",
			i,
			&vectorcito[i],
		)
	}
	/** Otra versión.
	* Analizar la salida. ¿Porqué da diferente que con &vectorcito[i]?
	 */
	fmt.Println()
	for i, v := range vectorcito {
		fmt.Printf(
			"La dirección de memoria del elemento %d de vectorcito es %d\n", i, &v)
	}
	/**
	  con slices
	*/
	tajadita := vectorcito[2:5]
	fmt.Println("\ntajadita: ", tajadita)
	fmt.Printf("La direccion de memoria de la tajadita es: %p\n", &tajadita)
	// cada posición del slice:
	for i := range tajadita {
		fmt.Printf(
			"La direccion de memoria del elemento %d de la tajadita es: %p\n", i,
			&tajadita[i])
	}
}