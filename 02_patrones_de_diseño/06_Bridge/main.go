package main

import (
	"bridge/parlantes"
	"bridge/reproductor"
	"fmt"
)

func main() {

	lista := []string{"Cancion 1", "Cancion 2", "Cancion 3", "Cancion 4"}

	parlanteBluetooth := parlantes.ParlanteBluetooth{}
	parlantePorCable := parlantes.ParlantePorCable{}

	reproductorMP3 := reproductor.MP3Player{}
	reproductorFLAC := reproductor.FLACPlayer{}

	fmt.Println("\n -------------- bluetooth / MP3 --------------")
	parlanteBluetooth.ConectarReproductor(&reproductorMP3)
	parlanteBluetooth.CargarLista(lista)

	parlanteBluetooth.Play()
	parlanteBluetooth.AdjustarVolume(23)
	parlanteBluetooth.Pausa()
	parlanteBluetooth.SiguienteCancion()
	parlanteBluetooth.Play()

	fmt.Println("\n -------------- bluetooth / FLAC --------------")
	parlanteBluetooth.ConectarReproductor(&reproductorFLAC)
	parlanteBluetooth.CargarLista(lista)
	parlanteBluetooth.Play()

	fmt.Println("\n -------------- cable / MP3 --------------")
	parlantePorCable.ConectarReproductor(&reproductorMP3)
	parlantePorCable.CargarLista(lista)
	parlantePorCable.Play()

}
