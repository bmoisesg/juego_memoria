package juego

import (
	"fmt"
)

func printMargen() string {
	for i := 0; i < margen; i++ {
		fmt.Print(" ")
	}
	return ""
}

//linealizar las posiciones x,y
func index(x int, y int) int {
	if x == 0 {
		return y
	} else if x == 1 {
		return filas + y
	} else {
		return (x)*filas + y
	}
}

func SetTarjetas(array []Tarjeta) {
	arrayTarjetas = array
}

func revisarTarjetasLevantadas() {
	var tarjeta1_x int = -1
	var tarjeta1_y int = -1
	var tarjeta2_x int = -1
	var tarjeta2_y int = -1

	for i := 0; i < len(arrayTarjetas); i++ {
		tarjeta := arrayTarjetas[i]
		if tarjeta.Mostrar == 1 && tarjeta.Delete == 0 {
			if tarjeta1_x == -1 {
				tarjeta1_x = tarjeta.X
				tarjeta1_y = tarjeta.Y
			} else {
				tarjeta2_x = tarjeta.X
				tarjeta2_y = tarjeta.Y
				break
			}
		}
	}

	if arrayTarjetas[index(tarjeta1_x, tarjeta1_y)].Number == arrayTarjetas[index(tarjeta2_x, tarjeta2_y)].Number {
		eliminarTarjeta(tarjeta1_x, tarjeta1_y)
		eliminarTarjeta(tarjeta2_x, tarjeta2_y)
		mensaje_si_pareja()

	} else {
		bajarTarjeta(tarjeta1_x, tarjeta1_y)
		bajarTarjeta(tarjeta2_x, tarjeta2_y)
		mensaje_no_son_pareja()
	}

	MostrarTarjetas()
}
