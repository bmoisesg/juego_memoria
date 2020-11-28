package juego

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func estadoTarjetaDelete(x int, y int) byte {
	return arrayTarjetas[index(x, y)].Delete
}

func pedirCoordenadas(mensaje string) (int, int) {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("-> ", mensaje, "tarjeta")
		fmt.Print("Ingresar coordenada x : ")
		coordenadaX, _ := reader.ReadString('\n')
		coordenadaX = strings.TrimSpace(coordenadaX)
		if coordenadaX == "adminadmin" {
			MostarTodasTarjetasVolteadas()
			pausa()
			continue
		}
		//verificarEntrada(coordenadaX)
		fmt.Print("Ingresar coordenada y : ")
		coordenadaY, _ := reader.ReadString('\n')
		coordenadaY = strings.TrimSpace(coordenadaY)
		x, error1 := strconv.Atoi(coordenadaX)
		y, error2 := strconv.Atoi(coordenadaY)
		//fmt.Println(error1)
		//fmt.Println(error2)
		if error1 == nil && error2 == nil {
			//fmt.Println("->", y, columnas)
			if x < 6 && y < filas {
				if estadoTarjetaDelete(x, y) == 0 {
					voltearTarjeta(x, y)
					MostrarTarjetas()
					//fmt.Println("(", x, ",", y, ")")
					return x, y
				} else {
					fmt.Println()
					fmt.Println("Esta tarjeta ya esta eliminaada")
				}
			}
		}
		fmt.Println()
		fmt.Println("Error: coordenadas no validas, ingresar nuevamente...")
		pausa()
		MostrarTarjetas()
	}
}

func IteracionJugador() bool {
	tarjeta1_x, tarjeta1_y := pedirCoordenadas("Primera")
	tarjeta2_x, tarjeta2_y := pedirCoordenadas("Segunda")

	if tarjeta1_x == tarjeta2_x && tarjeta1_y == tarjeta2_y {
		fmt.Println()
		fmt.Println("Error: Ingresaste la misma tarjeta")
		bajarTarjeta(tarjeta1_x, tarjeta1_y)
		pausa()
		MostrarTarjetas()
	} else {
		revisarTarjetasLevantadas()
	}
	//revisar que todas las tarjetas esten eliminadas
	return verificarTarjetas()

}

func verificarTarjetas() bool {
	for i := 0; i < len(arrayTarjetas); i++ {
		if arrayTarjetas[i].Delete == 0 {
			return false
		}
	}
	mensaje_you_win()
	return true
}
