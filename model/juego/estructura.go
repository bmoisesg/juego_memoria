package juego

type Tarjeta struct {
	Mostrar byte
	X       int
	Y       int
	Number  int
	Delete  byte
}

var margen = 30
var filas = 0
var columnas = 6
var arrayTarjetas []Tarjeta
