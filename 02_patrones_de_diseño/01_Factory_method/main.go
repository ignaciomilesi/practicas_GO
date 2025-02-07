package main

import (
	canal "Factory_method/canalDeComunicacion"
	"fmt"
)

func infoDeEnvio(c canal.Canal) {

	fmt.Printf("Canal de envio: %s. Remitente: %s.\n", c.GetMedio(), c.GetRemitente())

}

func main() {

	sms, _ := canal.Factory("SMS")
	email, _ := canal.Factory("Email")

	fmt.Println(sms.GetRemitente())
	sms.EnviarMensaje("Hola por SMS", "Ignacio")
	email.EnviarMensaje("Hola por Email", "Ignacio2")

	infoDeEnvio(sms)
	infoDeEnvio(email)

	sms2, _ := canal.Factory("SMS")

	sms2.EnviarMensaje("Hola 2 por SMS", "El otro Ignacio")

	infoDeEnvio(sms2)

	carta, _ := canal.Factory("Carta")

	carta.EnviarMensaje("Hojas en el viento", "Yo")

	//verifico que carta (debido a la interface es manejado como Canal) es del tipo Carta
	canal.SeleccionarTipoCarta(carta.(*canal.Carta), "Telegrama")

	carta.EnviarMensaje("Hojas en el viento", "Yo")

}
