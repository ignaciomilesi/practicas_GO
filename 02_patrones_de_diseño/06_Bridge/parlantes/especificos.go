package parlantes

import (
	"bridge/reproductor"
	"fmt"
)

type ParlanteBluetooth struct {
	generico
}

type ParlantePorCable struct {
	generico
}

func (b *ParlanteBluetooth) ConectarReproductor(nuevoReproductor reproductor.Reproductor) {
	b.generico.ConectarReproductor(nuevoReproductor)
	fmt.Printf("Reproductor %s conectado a parlantes Bluetooth\n", nuevoReproductor.Nombre())
}

func (c *ParlantePorCable) ConectarReproductor(nuevoReproductor reproductor.Reproductor) {
	c.generico.ConectarReproductor(nuevoReproductor)
	fmt.Printf("Reproductor %s conectado a parlantes por cable\n", nuevoReproductor.Nombre())
}
