package main

import "fmt"

type Bicicleta interface {
	Frenar()
}

// LEGACY
// Implementación de la bicicleta playera
type Playera struct{}

// LEGACY
// Frenar la bicicleta playera
func (p *Playera) PedalearHaciaAtras() {
	fmt.Println("Frenando la bicicleta playera...")
}

// Implementación de la bicicleta de carrera
type BicicletaDeCarreras struct{}

// Frenar la bicicleta de carrera
func (c *BicicletaDeCarreras) Frenar() {
	fmt.Println("Frenando la bicicleta de carrera...")
}

// Adaptador de freno que permite utilizar el freno de la bicicleta playera como si fuera una bicicleta de carrera
type PlayeraAdapter struct {
	Playera
}

// Frenar la bicicleta playera como si fuera una bicicleta de carrera
func (b *PlayeraAdapter) Frenar() {
	fmt.Println("Adaptando el freno de la bicicleta playera para el robot...")
	fmt.Print(" -> ")
	b.PedalearHaciaAtras()
}

func main() {
	// Instanciar bicicleta playera y de carrera
	playera := &Playera{}
	carreras := &BicicletaDeCarreras{}
	// Frenar la bicicleta playera y de carrera utilizando sus propios frenos
	playera.PedalearHaciaAtras()
	carreras.Frenar()
	fmt.Println("--------------------")
	bicicletas := []Bicicleta{&BicicletaDeCarreras{}, &PlayeraAdapter{}}
	bicicletas[0].Frenar()
	bicicletas[1].Frenar()
	fmt.Println("--------------------")
	// Utilizar el adaptador de freno para frenar la bicicleta playera como si fuera una bicicleta de carrera
	adapter := &PlayeraAdapter{}
	adapter.Frenar()
}