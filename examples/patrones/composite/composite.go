package main

import "fmt"

// La interfaz Componente define la estructura básica de los elementos
// que pueden ser hojas o compuestos del árbol.
type Componente interface {
	ObtenerNombre() string
	ObtenerCantidadArchivos() int
}

// El tipo Archivo representa una hoja del árbol. Contiene el nombre y la ruta del archivo.
type Archivo struct {
	nombre string
}

// El método ObtenerNombre devuelve el nombre del archivo.
func (a *Archivo) ObtenerNombre() string {
	return a.nombre
}

// El método ObtenerCantidadArchivos devuelve 1, ya que un archivo es una hoja.
func (a *Archivo) ObtenerCantidadArchivos() int {
	return 1
}

// El tipo Directorio representa un componente compuesto del árbol.
type Directorio struct {
	nombre    string
	elementos []Componente
}

// El método ObtenerNombre devuelve el nombre del directorio.
func (d *Directorio) ObtenerNombre() string {
	return d.nombre
}

// El método ObtenerCantidadArchivos recorre todos los elementos del directorio
// y sus subdirectorios para contar la cantidad de archivos.
func (d *Directorio) ObtenerCantidadArchivos() int {
	cantidadArchivos := 0
	for _, elemento := range d.elementos {
		cantidadArchivos += elemento.ObtenerCantidadArchivos()
	}
	return cantidadArchivos
}

// El método AgregarElemento agrega un nuevo elemento al directorio.
func (d *Directorio) AgregarElemento(elemento Componente) {
	d.elementos = append(d.elementos, elemento)
}

func main() {
	principal := &Directorio{nombre: "dir111"}
	principal.AgregarElemento(&Archivo{nombre: "archivoA"})
	principal.AgregarElemento(&Archivo{nombre: "archivoB"})
	dir2 := &Directorio{nombre: "dir222"}
	dir2.AgregarElemento(&Archivo{nombre: "archivoC"})
	dir2.AgregarElemento(&Archivo{nombre: "archivoD"})
	principal.AgregarElemento(dir2)
	dir3 := &Directorio{nombre: "dir333"}
	dir3.AgregarElemento(&Archivo{nombre: "archivoE"})
	dir2.AgregarElemento(dir3)
	fmt.Println(principal.ObtenerNombre(), "contiene", principal.ObtenerCantidadArchivos(), "archivos")
}