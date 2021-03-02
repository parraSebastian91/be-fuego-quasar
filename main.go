package main

import (
	"fmt"
	"math"
)

type coordenada struct {
	x, y float64
}

type circulo struct {
	puntoCentro coordenada
	radio       float64
}

type satellite struct {
	name     string
	distance float64
	massage  []string
}

type satellites struct {
	satellites []satellite
}

func getDistanciaEntreCirculos(c1 circulo, c2 circulo) float64 {
	dxAB, dyAB := c2.puntoCentro.x-c1.puntoCentro.x, c2.puntoCentro.y-c1.puntoCentro.y
	dAB := math.Sqrt(dxAB*dxAB + dyAB*dyAB)
	return dAB
}

func (c *satellites) intersecion_entre_circunferencias(cA circulo, cB circulo) (coordenada, coordenada) {

	dxAB, dyAB := cB.puntoCentro.x-cA.puntoCentro.x, cB.puntoCentro.y-cA.puntoCentro.y

	distanciaAB := getDistanciaEntreCirculos(cA, cB)

	if distanciaAB > cA.radio+cB.radio {
		// circulos separados, no se puede determinar
	}

	if distanciaAB < math.Abs(cA.radio-cB.radio) {

	}

	if distanciaAB == 0 && cA.radio == cB.radio {

	}

	a := (cA.radio*cA.radio - cB.radio*cB.radio + distanciaAB*distanciaAB) / (2 * distanciaAB)
	h := math.Sqrt(cA.radio*cA.radio - a*a)
	xm := cA.puntoCentro.x + a*dxAB/distanciaAB
	ym := cA.puntoCentro.y + a*dyAB/distanciaAB
	xs1 := xm + h*dyAB/distanciaAB
	xs2 := xm - h*dyAB/distanciaAB
	ys1 := ym - h*dxAB/distanciaAB
	ys2 := ym + h*dxAB/distanciaAB
	return coordenada{x: xs1, y: ys1}, coordenada{x: xs2, y: ys2}
}

func main() {
	s := satellites{
		satellites: []satellite{
			satellite{
				name:     "kenobi",
				distance: 485.69,
				massage:  []string{"este", "", "", "mensaje", ""},
			}, satellite{
				name:     "skywalker",
				distance: 266.08,
				massage:  []string{"", "es", "", "", "secreto"},
			}, satellite{
				name:     "sato",
				distance: 600.5,
				massage:  []string{"este", "", "un", "", ""},
			},
		},
	}

	kenobi := circulo{
		puntoCentro: coordenada{x: -500, y: -200},
		radio:       s.satellites[0].distance,
	}
	skywalker := circulo{
		puntoCentro: coordenada{x: 100, y: -100},
		radio:       s.satellites[1].distance,
	}
	sato := circulo{
		puntoCentro: coordenada{x: 500, y: 100},
		radio:       s.satellites[2].distance,
	}

	punto1, punto2 := s.intersecion_entre_circunferencias(kenobi, skywalker)
	punto3, punto4 := s.intersecion_entre_circunferencias(skywalker, sato)
	punto5, punto6 := s.intersecion_entre_circunferencias(kenobi, sato)
	fmt.Println("punto 1: (", math.Round(punto1.x), ",", math.Round(punto1.y), ")")
	fmt.Println("punto2: (", math.Round(punto2.x), ",", math.Round(punto2.y), ")")
	fmt.Println("punto 3: (", math.Round(punto3.x), ",", math.Round(punto3.y), ")")
	fmt.Println("punto4: (", math.Round(punto4.x), ",", math.Round(punto4.y), ")")
	fmt.Println("punto 5: (", math.Round(punto5.x), ",", math.Round(punto5.y), ")")
	fmt.Println("punto6: (", math.Round(punto6.x), ",", math.Round(punto6.y), ")")

}
