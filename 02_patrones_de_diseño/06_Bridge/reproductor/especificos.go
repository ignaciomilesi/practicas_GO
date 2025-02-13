package reproductor

import "fmt"

// MP3Player,
type MP3Player struct {
	generico
}

func (mp3 *MP3Player) CargarLista(nuevaLista []string) {
	mp3.generico.CargarLista(nuevaLista)
	fmt.Printf("el reproductor de MP3\n")
}

func (mp3 MP3Player) Nombre() string {
	return "MP3"
}

// FLACPlayer
type FLACPlayer struct {
	generico
}

func (flac *FLACPlayer) CargarLista(nuevaLista []string) {
	flac.generico.CargarLista(nuevaLista)
	fmt.Printf("el reproductor de FLAC\n")
}

func (flac FLACPlayer) Nombre() string {
	return "FLAC"
}
