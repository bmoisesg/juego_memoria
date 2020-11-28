package juego

import (
	"fmt"
	"strconv"
)

func MostrarTarjetas() {
	for i := 0; i < 30; i++ {
		fmt.Println()
	}

	fmt.Println(printMargen() + "y\\x 0    1    2    3    4    5    ")
	y := 0
	for j := 0; j < len(arrayTarjetas); j = j + 6 {
		fmt.Println(printMargen() + "  ╔═══╗╔═══╗╔═══╗╔═══╗╔═══╗╔═══╗")
		fmt.Print(printMargen() + strconv.Itoa(y) + " ")
		for i := 0; i < 6; i++ {
			tarjeta := arrayTarjetas[index(i, y)]
			if tarjeta.Delete == 1 {
				fmt.Print("║   ║")
			} else if tarjeta.Mostrar == 1 {
				if tarjeta.Number < 10 {
					fmt.Print("║ " + strconv.Itoa(tarjeta.Number) + " ║")
				} else {
					fmt.Print("║ " + strconv.Itoa(tarjeta.Number) + "║")
				}
			} else {
				fmt.Print("║ ░ ║")
			}
		}
		fmt.Println()
		fmt.Println(printMargen() + "  ╚═══╝╚═══╝╚═══╝╚═══╝╚═══╝╚═══╝")
		y++
	}
}

func MostarTodasTarjetasVolteadas() {
	fmt.Println("     0    1    2    3    4    5 ")
	y := 0
	for j := 0; j < len(arrayTarjetas); j = j + 6 {
		fmt.Println("   ╔═══╗╔═══╗╔═══╗╔═══╗╔═══╗╔═══╗")
		fmt.Print(" " + strconv.Itoa(y) + " ")
		for i := 0; i < 6; i++ {
			tarjeta := arrayTarjetas[index(i, y)]
			fmt.Print("║ ")
			numero := tarjeta.Number
			if numero < 10 {
				fmt.Print(strconv.Itoa(numero) + " ║")
			} else {
				fmt.Print(strconv.Itoa(numero) + "║")
			}
		}
		fmt.Println("\n   ╚═══╝╚═══╝╚═══╝╚═══╝╚═══╝╚═══╝")
		y++
	}
}
