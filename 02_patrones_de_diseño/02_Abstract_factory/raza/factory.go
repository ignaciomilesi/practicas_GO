package raza

import (
	"fmt"
)

func Factory(razaSeleccionada string) (raza, error) {

	switch razaSeleccionada {

	case "Elfo":
		return Elfo{}, nil

	case "Enano":
		return Enano{}, nil

	case "Orco":
		return Orco{}, nil

	default:
		return nil, fmt.Errorf("Raza inexistente")
	}
}
