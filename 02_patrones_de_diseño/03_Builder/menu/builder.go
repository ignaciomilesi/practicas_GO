package menu

import "fmt"

type Builder struct {
	mesa
}

func (b *Builder) AgregarEtapa(etapa, plato string) {

	switch etapa {

	case "entrada":
		b.entrada = plato

	case "principal":
		b.principal = plato

	case "bebida":
		b.bebida = plato

	case "postre":
		b.postre = plato

	default:
		fmt.Println("Selección equivocada")
	}
}

func (b Builder) Construir() Menu {

	var nuevoMenu Menu

	if b.entrada == "" {
		nuevoMenu.entrada = "No solicitado"
	} else {
		nuevoMenu.entrada = b.entrada
	}

	if b.principal == "" {
		nuevoMenu.principal = "No solicitado"
	} else {
		nuevoMenu.principal = b.principal
	}

	if b.bebida == "" {
		nuevoMenu.bebida = "No solicitado"
	} else {
		nuevoMenu.bebida = b.bebida
	}

	if b.postre == "" {
		nuevoMenu.postre = "No solicitado"
	} else {
		nuevoMenu.postre = b.postre
	}

	return nuevoMenu
}

func GetMenuItaliano() Menu {
	nuevoMenuItaliano := Menu{
		mesa{
			entrada:   "Bruschetta",
			principal: "Spaghetti",
			bebida:    "Limoncello",
			postre:    "Tiramisu",
		},
	}

	return nuevoMenuItaliano
}
