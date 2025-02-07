package canalDeComunicacion

import "fmt"

type Canal interface {
	EnviarMensaje(mensaje, remitente string)
	GetMedio() string
	GetRemitente() string
}

func Factory(notificationType string) (Canal, error) {

	switch notificationType {

	case "SMS":
		return CrearCanalSMS(), nil

	case "Email":
		return CrearCanalEmail(), nil

	case "Carta":
		return CrearCanalCarta(), nil

	default:
		return nil, fmt.Errorf("No Notification Type")
	}
}
