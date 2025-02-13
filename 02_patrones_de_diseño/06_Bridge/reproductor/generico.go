package reproductor

import "fmt"

type Reproductor interface {
	CargarLista(nuevaLista []string)
	Play()
	Pausa()
	SiguienteCancion()
	Nombre() string
}

type generico struct {
	listaCanciones []string
	seleccion      int
}

func (g *generico) CargarLista(nuevaLista []string) {
	g.listaCanciones = nuevaLista
	fmt.Printf("Se ha cargado la nueva lista en ")
}

func (g generico) Play() {
	fmt.Println("Se reproduce la canción:", g.listaCanciones[g.seleccion])
}

func (g generico) Pausa() {
	fmt.Println("Reproducción en pausa")
}

func (g *generico) SiguienteCancion() {

	g.seleccion = (g.seleccion + 1) % len(g.listaCanciones)
	fmt.Println("Se cambio a la canción:", g.listaCanciones[g.seleccion])
}
