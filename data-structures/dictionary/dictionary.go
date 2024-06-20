// Package dictionary proporciona una implementación de un diccionario genérico basado en un map de Go.
// Expone la estructura Diccionario y sus métodos para manipular un diccionario
package dictionary

import (
	"errors"
	"fmt"
)

// Diccionario implementa un diccionario genérico basado en un map de Go.
// Las claves deben ser de un tipo comparable y los valores pueden ser de cualquier tipo.
type Diccionario[Clave comparable, Valor any] struct {
	diccionario map[Clave]Valor
}

// NuevoDiccionario crea un nuevo diccionario vacío.
//
// Uso:
//
//	Diccionario := dictionary.NuevoDiccionario[int, string]() // Crea un diccionario de enteros a cadenas.
func NuevoDiccionario[Clave comparable, Valor any]() *Diccionario[Clave, Valor] {
	return &Diccionario[Clave, Valor]{diccionario: make(map[Clave]Valor)}
}

// Insertar inserta el par (clave, valor) en el Diccionario. Si la `clave` existe, reemplaza
// el par existente; si no existe, el par es agregado al diccionario.
//
// Uso:
//
//	Diccionario.Insertar(10, "ten") // Agrega el par (10, "ten") al diccionario.
//
// Parámetros:
//
//	`clave`: la clave del par a insertar.
//	`valor`: el valor del par a insertar.
func (Diccionario *Diccionario[Clave, Valor]) Insertar(clave Clave, valor Valor) {
	Diccionario.diccionario[clave] = valor
}

// Contiene verifica si la `clave` especificada existe en el diccionario
//
// Uso:
//
//	if Diccionario.Contiene(10) {
//		fmt.Println("Clave 10 existe.")
//	}
//
// Parámetros:
//
//	`clave`: la clave a buscar en el diccionario.
//
// Retorna:
//
//	`true` si la clave existe en el diccionario; `false` en caso contrario.
func (Diccionario *Diccionario[Clave, Valor]) Contiene(clave Clave) bool {
	_, existe := Diccionario.diccionario[clave]
	return existe
}

// Obtener devuelve el `valor` asociado a la `clave` específicada.
//
// Uso:
//
//	if valor, err := Diccionario.Obtener(10); err != nil {
//		fmt.Println("Valor asociado a la clave 10:", valor)
//	}
//
// Parámetros:
//
//	`clave`: la clave cuyo valor asociado se desea obtener.
//
// Retorna:
//
//	El valor asociado a la clave especificada  y nil como error, o el nil type del tipo Valor
//	y un error si la clave no existe.
func (Diccionario *Diccionario[Clave, Valor]) Obtener(clave Clave) (Valor, error) {
	valor, existe := Diccionario.diccionario[clave]
	if !existe {
		return valor, errors.New("clave inexistente")
	}
	return valor, nil
}

// Remover remueve el par asociado a la `clave` especificada del diccionario.
//
// Uso:
//
//	Diccionario.Remover(10)
//
// Parámetros:
//
//	`clave`: la clave a remover del diccionario.
func (Diccionario *Diccionario[Clave, Valor]) Remover(clave Clave) {
	delete(Diccionario.diccionario, clave)
}

// Tamaño devuelve el tamaño del diccionario.
//
// Uso:
//
//	fmt.Println("Tamaño del diccionario:", Diccionario.Tamaño())
//
// Retorna:
//
//	La cantidad de entradas presentes en el diccionario.
func (Diccionario *Diccionario[Clave, Valor]) Tamaño() int {
	return len(Diccionario.diccionario)
}

// Claves devuelve todas las claves del diccionario. Por la naturaleza de los map de Go,
// las claves no se devuelven en un orden específico. No puede contener duplicados.
//
// Uso:
//
//	keys := Diccionario.Claves()
//
// Retorna:
//
//	Un slice de todas las claves presentes en el diccionario.
func (Diccionario *Diccionario[Clave, Valor]) Claves() []Clave {
	clavesDiccionario := make([]Clave, 0, Diccionario.Tamaño())
	for clave := range Diccionario.diccionario {
		clavesDiccionario = append(clavesDiccionario, clave)
	}
	return clavesDiccionario
}

// Valores devuelve todos los valores del diccionario. Por la naturaleza de los map de Go,
// los valores no se devuelven en un orden específico. Puede contener duplicados.
//
// Uso:
//
//	values := Diccionario.Valores()
//
// Retorna:
//
//	Un slice de todos los valores presentes en el diccionario.
func (Diccionario *Diccionario[Clave, Valor]) Valores() []Valor {
	valoresDiccionario := make([]Valor, 0, Diccionario.Tamaño())
	for _, valor := range Diccionario.diccionario {
		valoresDiccionario = append(valoresDiccionario, valor)
	}
	return valoresDiccionario
}

// String devuelve la representación en `string` del diccionario.
//
// Uso:
//
//	fmt.Println(Diccionario)
//
// Retorna:
//
//	Una representación en `string` del diccionario.
func (Diccionario *Diccionario[Clave, Valor]) String() string {
	if Diccionario.Tamaño() == 0 {
		return "Diccionario: {}"
	}
	String := "Diccionario: {\n"
	for clave, valor := range Diccionario.diccionario {
		String += fmt.Sprintf("  %v: %v\n", clave, valor)
	}
	String += "}"
	return String
}