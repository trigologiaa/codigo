package composite

type Rectangulo struct {
	base   float64
	altura float64
}

func NewRectangulo(base, altura float64) Rectangulo {
	return Rectangulo{base: base, altura: altura}
}

func (r Rectangulo) Area() float64 {
	return r.altura * r.base
}