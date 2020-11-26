package main

import (
	"bufio"
	"fmt"
	"juegoMemoria/model/jugador"
	"os"
)

func main() {
	fmt.Println("Bienvenido al juego de memoria")
	fmt.Print("Ingresar nombre:")
	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n')
	jugador1 := jugador.USER{
		Life:  100,
		Score: 0,
		Name:  name}
	fmt.Println(jugador1)
	jugador.Hola()
}
