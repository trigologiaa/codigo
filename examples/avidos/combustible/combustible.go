package combustible

// MinimasRecargas recibe la distancia a recorrer, la capacidad del tanque y las estaciones de servicio
// y devuelve la cantidad mínima de recargas necesarias para recorrer la distancia.
func MinimasRecargas(distancia int, capacidad int, estaciones []int) int {
	return minimasRecargas(0, distancia, capacidad, estaciones)
}

func minimasRecargas(ubicacion int, distancia int, capacidad int, estaciones []int) int {
	// alcanza el tanque para llegar al final
	if ubicacion+capacidad >= distancia {
		return 0
	}
	ultimaParada := ubicacion
	// IDEA: me fijo hasta dónde pueda llegar desde donde estoy ahora
	// mientras haya estaciones, y tenga la capacidad de llegar a la próxima
	for len(estaciones) > 0 && capacidad >= estaciones[0]-ubicacion {
		// me movería a la próxima estación
		ultimaParada = estaciones[0]
		// descarto la estación donde no necesito parar
		estaciones = estaciones[1:]
	}
	return 1 + minimasRecargas(ultimaParada, distancia, capacidad, estaciones)
}