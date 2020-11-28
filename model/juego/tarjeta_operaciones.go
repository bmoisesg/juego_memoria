package juego

func bajarTarjeta(x int, y int) {
	arrayTarjetas[index(x, y)].Mostrar = 0
}

func eliminarTarjeta(x int, y int) {
	arrayTarjetas[index(x, y)].Mostrar = 0
	arrayTarjetas[index(x, y)].Delete = 1
}

func voltearTarjeta(x int, y int) {
	arrayTarjetas[index(x, y)].Mostrar = 1
}
