package composite

type FiguraCompuesta struct {
	figuras []Figura
}

func NewFiguraCompuesta() FiguraCompuesta {
	return FiguraCompuesta{}
}

func (fc FiguraCompuesta) Area() float64 {
	var area float64
	for _, v := range fc.figuras {
		area += v.Area()
	}
	return area
}

func (fc *FiguraCompuesta) Add(f Figura) {
	fc.figuras = append(fc.figuras, f)
}