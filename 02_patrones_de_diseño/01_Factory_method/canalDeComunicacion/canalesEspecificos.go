package canalDeComunicacion

import "fmt"

//////////////////////////////// SMS ////////////////////////////////

type SMS struct {
	Generico
}

func CrearCanalSMS() Canal {
	return &SMS{
		Generico: Generico{
			medio: "SMS",
		},
	}
}

//////////////////////////////// Email ////////////////////////////////

type Email struct {
	Generico
}

func CrearCanalEmail() Canal {
	return &Email{
		Generico: Generico{
			medio: "Email",
		},
	}
}

//////////////////////////////// Carta ////////////////////////////////

type Carta struct {
	Generico
	tipo string
}

func (c *Carta) EnviarMensaje(mensaje, remitente string) {
	c.remitente = remitente
	fmt.Printf("Mensaje enviado: %s. Tipo de carta: %s \n", mensaje, c.tipo)
}

func SeleccionarTipoCarta(carta *Carta, tipo string) {
	carta.tipo = tipo
}

func CrearCanalCarta() Canal {
	return &Carta{
		Generico: Generico{
			medio: "Carta",
		},
		tipo: "Correo standard",
	}
}
