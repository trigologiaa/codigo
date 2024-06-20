package composite

type Triangulo struct {
	base   float64
	altura float64
}

func NewTriangulo(base, altura float64) Triangulo {
	return Triangulo{base: base, altura: altura}
}

func (r Triangulo) Area() float64 {
	return r.altura * r.base / 2
}