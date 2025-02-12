package main

import (
	"adapter/agujero"
	"adapter/pieza"
	"fmt"
)

func main() {

	piezaRed := pieza.Redonda{}
	agujeroRed := agujero.Redondo{}

	piezaRed.SetRadio(10)
	agujeroRed.SetRadio(15)

	fmt.Printf("Encaje Red-Red: %t \n", agujeroRed.ComprobarEncaje(piezaRed))

	piezaCuad := pieza.Cuadrada{}

	piezaCuad.SetLado(10)

	fmt.Printf(
		"Encaje Red-Cuad: %t \n",
		agujeroRed.ComprobarEncaje(pieza.AdaptadorRedondoCuadrado(piezaCuad)))

}
