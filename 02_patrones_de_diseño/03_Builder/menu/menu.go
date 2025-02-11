package menu

import "fmt"

type mesa struct {
	entrada   string
	principal string
	bebida    string
	postre    string
}

type Menu struct {
	mesa
}

func (m *Menu) ConsumirEtapa(etapa string) {
	switch etapa {
	case "entrada":
		if m.entrada != "No solicitado" {
			m.entrada = "Consumido"
		}
	case "principal":
		if m.principal != "No solicitado" {
			m.principal = "Consumido"
		}
	case "bebida":
		if m.bebida != "No solicitado" {
			m.bebida = "Consumido"
		}
	case "postre":
		if m.postre != "No solicitado" {
			m.postre = "Consumido"
		}
	default:
		fmt.Println("Selección equivocada")
	}
}

func (m Menu) EmitirResumen() {
	fmt.Println("\nResumen de la mesa:")
	fmt.Println(" - Entrada:", m.entrada)
	fmt.Println(" - Principal:", m.principal)
	fmt.Println(" - Bebida:", m.bebida)
	fmt.Println(" - Postre:", m.postre)
}
