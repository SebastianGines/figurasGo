package figuras

import "math"

type Circulo struct {
	Radio float32
}

func (circulo *Circulo) area() float32 {
	return math.Pi * (circulo.Radio * circulo.Radio)
}

func (circulo *Circulo) perimetro() float32 {
	return 2 * math.Pi * circulo.Radio
}
