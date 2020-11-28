package main

import (
	"bufio"
	"fmt"
	"juegoMemoria/model/juego"
	"os"
	"strconv"
	"strings"
)

func main() {
	for {
		saltos()
		juego.Mensaje_bienvenidas()
		opcion := getOpcion()
		flagSalir := false
		switch opcion {
		case 1:
			tarjetas := juego.CrearTablero(1)
			for {
				saltos()
				juego.Mensaje_seleccionar_nivel()
				opcion := getOpcion()
				if opcion == 1 {
					break
				} else if opcion == 2 {
					tarjetas = juego.CrearTablero(2)
					break
				} else if opcion == 3 {
					tarjetas = juego.CrearTablero(3)
					break
				}

			}
			//tarjetas := juego.CrearTablero(1)
			juego.SetTarjetas(tarjetas)
			juego.MostarTodasTarjetasVolteadas()
			juego.MostrarTarjetas()

			for {
				condicion := juego.IteracionJugador()
				if condicion {
					break
				}
			}
			break
		case 2:
			saltos()
			juego.Mensaje_Acerda_de()
			break
		case 3:
			flagSalir = true
			break
		}
		if flagSalir {
			saltos()
			fmt.Println("Saliendo del juego, que tenga un buen dia (:")
			break
		}
	}

}
func getOpcion() int {
	for {
		fmt.Print("Ingrese opcion: ")
		reader := bufio.NewReader(os.Stdin)
		data, _ := reader.ReadString('\n')
		data = strings.TrimSpace(data)
		x, error1 := strconv.Atoi(data)
		if error1 == nil {
			return x
		}
		fmt.Println("Error: no valido, vuelva a ingresar datos")
	}
}

func saltos() {
	for i := 0; i < 100; i++ {
		fmt.Println()
	}
}
