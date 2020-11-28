package juego

import (
	"math/rand"
	"time"
)

func CrearTablero(nivel byte) []Tarjeta {
	arreglo := []Tarjeta{}
	array_parejas := []int{}
	rand.Seed(time.Now().Unix())

	switch nivel {
	case 1:
		filas = 2
		array_parejas = generarParejas(filas * columnas) // 6 columnas del juego
		break
	case 2:
		filas = 3
		array_parejas = generarParejas(filas * columnas)
		break
	case 3:
		filas = 5
		array_parejas = generarParejas(filas * columnas)
		break
	}

	var z int = 0
	for i := 0; i < 6; i++ {
		for j := 0; j < filas; j++ {
			arreglo = append(arreglo, Tarjeta{0, i, j, array_parejas[z], 0})
			z++
		}
	}
	return arreglo
}

//Almacena en un array de enteros las parejas desordenadas  [0,1,0,2,1,3,3,2...]
func generarParejas(posiciones int) []int {
	arreglo := []int{}
	for i := 0; i < posiciones; i++ {
		for {
			num := rand.Intn(posiciones / 2)
			apariciones := getVecesExiste(num, arreglo)
			if apariciones < 2 { //solo puede existir 2 tarjetas iguales
				arreglo = append(arreglo, num)
				break
			}
		}
	}
	return arreglo
}

//Revisa cuantas veces aparece una tarjeta
func getVecesExiste(num int, array []int) int {
	contador := 0
	for i := 0; i < len(array); i++ {
		if num == array[i] {
			contador++
		}
	}
	return contador
}
