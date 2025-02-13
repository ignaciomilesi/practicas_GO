package parlantes

import (
	"bridge/reproductor"
	"fmt"
)

type generico struct {
	reproductor reproductor.Reproductor
	volumen     int
}

func (g *generico) ConectarReproductor(nuevoReproductor reproductor.Reproductor) {
	g.reproductor = nuevoReproductor
}

func (g *generico) AdjustarVolume(nuevoVolumen int) {
	g.volumen = nuevoVolumen
	fmt.Println("Volumen ajustado a:", nuevoVolumen)
}

// func que lanzan a los func del reproductor

func (g *generico) CargarLista(nuevaLista []string) {
	g.reproductor.CargarLista(nuevaLista)
}

func (g generico) Play() {
	g.reproductor.Play()
}

func (g generico) Pausa() {
	g.reproductor.Pausa()
}

func (g generico) SiguienteCancion() {
	g.reproductor.SiguienteCancion()
}
