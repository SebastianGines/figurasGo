package figuras

type Cuadrado struct {
	Alto  float32
	Ancho float32
}

func (cuadrado *Cuadrado) area() float32 {
	return cuadrado.Alto * cuadrado.Ancho
}

func (cuadrado *Cuadrado) perimetro() float32 {
	return 2*cuadrado.Alto + 2*cuadrado.Ancho
}
