package figuras

import (
	"fmt"
	"reflect"
)

type Figura interface {
	area() float32
	perimetro() float32
}

func CalcularArea(figura Figura) {
	fmt.Println("El área de:", reflect.TypeOf(figura).Elem().Name(), "es:", figura.area())
}

func CalcularPerimetro(figura Figura) {
	fmt.Println("El perímetro de", reflect.TypeOf(figura).Elem().Name(), "es:", figura.perimetro())
}
