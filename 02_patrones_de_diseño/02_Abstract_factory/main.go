package main

import (
	"Abstract_Factory/raza"
	"fmt"
)

func main() {

	razas := []string{"Elfo", "Enano", "Orco"}

	for _, nombreRaza := range razas {
		fmt.Println("\n----", nombreRaza, "----")

		fabrica, err := raza.Factory(nombreRaza)

		if err != nil {
			fmt.Println(err)
			return
		}

		if arma, err := fabrica.GenerarArma(); err == nil {
			arma.Describir()
			arma.Usar()
			arma.Mejorar()
			arma.Usar()
		} else {
			fmt.Println(err)
		}

		fmt.Println("----")

		if armadura, err := fabrica.GenerarArmadura(); err == nil {
			armadura.Describir()
		} else {
			fmt.Println(err)
		}

		fmt.Println("----")

		if casco, err := fabrica.GenerarCasco(); err == nil {
			casco.Describir()
		} else {
			fmt.Println(err)
		}

	}

}
