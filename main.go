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

func main() {
	s := satellites{
		satellites: []satellite{
			{
				name:     "kenobi",
				distance: 485.69,
				massage:  []string{"este", "", "", "mensaje", ""},
			}, {
				name:     "skywalker",
				distance: 266.08,
				massage:  []string{"", "es", "", "", "secreto"},
			}, {
				name:     "sato",
				distance: 600.5,
				massage:  []string{"este", "", "un", "", ""},
			},
		},
	}
	fmt.Println(GetLocation(s))
}

//
func GetLocation(distances satellites) coordenada {

	var kenobi circulo
	var skywalker circulo
	var sato circulo
	for _, e := range distances.satellites {
		if e.name == "kenobi" {
			kenobi = circulo{
				puntoCentro: coordenada{x: -500, y: -200},
				radio:       e.distance,
			}
		} else if e.name == "skywalker" {
			skywalker = circulo{
				puntoCentro: coordenada{x: 100, y: -100},
				radio:       e.distance,
			}
		} else if e.name == "sato" {
			sato = circulo{
				puntoCentro: coordenada{x: 500, y: 100},
				radio:       e.distance,
			}
		}
	}

	punto1, punto2 := distances.intersecion_entre_circunferencias(kenobi, skywalker)
	punto3, punto4 := distances.intersecion_entre_circunferencias(skywalker, sato)
	punto5, punto6 := distances.intersecion_entre_circunferencias(kenobi, sato)
	arrayCoordenadas := []coordenada{punto1, punto2, punto3, punto4, punto5, punto6}
	return filtrarCoordenada(arrayCoordenadas)
}

func filtrarCoordenada(arr []coordenada) coordenada {
	dict := make(map[coordenada]int)
	var coorBuscada coordenada = arr[0]
	for _, punto := range arr {
		dict[punto] = dict[punto] + 1
	}
	for key, element := range dict {
		if dict[coorBuscada] < element {
			coorBuscada = key
		}
	}
	return coorBuscada
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
	return coordenada{x: math.Round(xs1*10) / 10, y: math.Round(ys1*10) / 10}, coordenada{x: math.Round(xs2*10) / 10, y: math.Round(ys2*10) / 10}
}
