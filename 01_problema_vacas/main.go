package main

import (
	"fmt"
	"problema_vacas/pkg/listadoVaca"
	"problema_vacas/pkg/selector"
	"problema_vacas/scripts"
)

func main() {

	prueba1 := "6 700 360,250,400,180,50,90 40,35,43,28,12,13"
	prueba2 := "8 1000 223,243,100,200,200,155,300,150 30,34,28,45,31,50,29,10"
	prueba3 := "10 2000 340,355,223,243,130,240,260,155,302,130 45,50,34,39,29,40,30,52,31,15"

	c := make(chan string)

	go realizarEleccion("Prueba 1", prueba1, c)
	go realizarEleccion("Prueba 2", prueba2, c)
	go realizarEleccion("Prueba 3", prueba3, c)

	fmt.Print(<-c)
	fmt.Print(<-c)
	fmt.Print(<-c)

}

func realizarEleccion(nombreEleccion, entrada string, c chan string) {

	txt := fmt.Sprintln("\n#################################", nombreEleccion, "#################################")

	limitPesoCamion, listaPeso, listaProduccion := scripts.ParseoEntrada(entrada)

	var listaVacasEnMercado listadoVaca.Lista

	listaVacasEnMercado.Crear(listaPeso, listaProduccion)

	selector := selector.Selector{
		ListadoVacaEnMercado: listaVacasEnMercado,
		LimiteCamion:         limitPesoCamion,
	}

	txt += fmt.Sprintln("Limite camion:", limitPesoCamion)

	txt += fmt.Sprintln("Listado de vacas:", listaVacasEnMercado.Lista)

	txt += selector.SelecionarVacas(true)

	c <- txt
}
