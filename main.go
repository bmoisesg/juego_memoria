package main

import (
	"fmt"

	"juegoMemoria/model/juego"
)

func main() {
	fmt.Println("Bienvenido al juego de memoria")
	tarjetas := juego.CrearTablero(1)
	juego.SetTarjetas(tarjetas)
	juego.MostarTodasTarjetasVolteadas()
	juego.MostrarTarjetas()
	for {
		juego.IteracionJugador()
	}
}
