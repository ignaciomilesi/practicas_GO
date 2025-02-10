package raza

import (
	"Abstract_Factory/equipamiento/arma"
	"Abstract_Factory/equipamiento/armadura"
)

type Orco struct {
	Generico
}

func (o Orco) GenerarArma() (arma.Arma, error) {
	return &arma.OrcoArma{
		Generica: arma.Generica{
			Nombre: "Garrote",
			Tipo:   "Corta distancia",
			Daño:   10,
		},
	}, nil
}

func (o Orco) GenerarArmadura() (armadura.Armadura, error) {
	return armadura.OrcoArmadura{
		Generica: armadura.Generica{
			Nombre:     "Pechera de acero",
			Tipo:       "pesada",
			Proteccion: 25,
		},
	}, nil
}
