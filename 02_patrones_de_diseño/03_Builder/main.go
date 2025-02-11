package main

import "builder/menu"

func main() {

	menuBuilder := menu.Builder{}

	menuBuilder.AgregarEtapa("entrada", "Provoleta")
	menuBuilder.AgregarEtapa("bebida", "Vino")
	menuBuilder.AgregarEtapa("principal", "asado")

	menu1 := menuBuilder.Construir()

	menu1.EmitirResumen()

	menu1.ConsumirEtapa("entrada")
	menu1.ConsumirEtapa("bebida")
	menu1.ConsumirEtapa("postre") //no hace nada ya que no fue agregado

	menu1.EmitirResumen()

	menuItaliano := menu.GetMenuItaliano()

	menuItaliano.EmitirResumen()

}
