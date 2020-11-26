package jugador

import "fmt"

type USER struct {
	Life  int8
	Score int8
	Name  string
}

func Hola() {
	fmt.Println("      1    2    3  ")
	fmt.Println("    ╔═══╗╔═══╗╔═══╗")
	fmt.Println(" 1  ║   ║║   ║║   ║")
	fmt.Println("    ╚═══╝╚═══╝╚═══╝")
	fmt.Println("    ╔═══╗╔═══╗╔═══╗")
	fmt.Println(" 2  ║   ║║   ║║   ║")
	fmt.Println("    ╚═══╝╚═══╝╚═══╝")
}
