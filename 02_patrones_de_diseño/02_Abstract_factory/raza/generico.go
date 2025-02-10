package raza

import (
	"Abstract_Factory/equipamiento/arma"
	"Abstract_Factory/equipamiento/armadura"
	"Abstract_Factory/equipamiento/casco"
	"fmt"
)

type raza interface {
	GenerarArma() (arma.Arma, error)
	GenerarArmadura() (armadura.Armadura, error)
	GenerarCasco() (casco.Casco, error)
}

type Generico struct {
}

func (g Generico) GenerarArma() (arma.Arma, error) {
	return nil, fmt.Errorf("Equipamiento no admitido para la raza seleccionada")
}

func (g Generico) GenerarArmadura() (armadura.Armadura, error) {
	return nil, fmt.Errorf("Equipamiento no admitido para la raza seleccionada")
}

func (g Generico) GenerarCasco() (casco.Casco, error) {
	return nil, fmt.Errorf("Equipamiento no admitido para la raza seleccionada")
}
