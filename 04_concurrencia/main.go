package main

import "concurrencia/ejemplo"

func main() {

	// Descomentar el ejemplo que se quiera ver

	//ejemplo.SinWaitGroup()
	//ejemplo.UsoWaitGroup()
	//ejemplo.CreandoChannel()
	//ejemplo.BloqueoChannelConSleep()
	//ejemplo.BloqueoChannelConOtroChannel()
	ejemplo.DeadlockPorDobleenvio()
}
