package composite

import "math"

type Circulo struct {
	radio float64
}

func NewCirculo(radio float64) Circulo {
	return Circulo{radio: radio}
}

func (c Circulo) Area() float64 {
	return c.radio * c.radio * math.Pi
}