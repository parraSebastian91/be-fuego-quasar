package geometria

import (
	"math"
)

type coordenada struct {
	x, y float32
}

type circulo struct {
	puntoCentro coordenada
	radio float32
}

func (c *circulo) intersecion_entre_circunferencias () (coordenada,coordenada)  {

}