package canalDeComunicacion

import "fmt"

type Generico struct {
	medio     string
	remitente string
}

func (g *Generico) EnviarMensaje(mensaje, remitente string) {
	g.remitente = remitente
	fmt.Println(mensaje)
}

func (g Generico) GetMedio() string {
	return g.medio
}

func (g Generico) GetRemitente() string {

	if g.remitente == "" {

		return "No se han enviado msj"
	}
	return g.remitente
}
