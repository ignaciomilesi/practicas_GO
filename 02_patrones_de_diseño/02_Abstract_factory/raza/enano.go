package raza

import (
	"Abstract_Factory/equipamiento/arma"
	"Abstract_Factory/equipamiento/armadura"
	"Abstract_Factory/equipamiento/casco"
)

type Enano struct {
	Generico
}

func (e Enano) GenerarArma() (arma.Arma, error) {
	return &arma.EnanoArma{
		Generica: arma.Generica{
			Nombre: "Hacha",
			Tipo:   "Corta distancia",
			Daño:   20,
		},
	}, nil
}

func (e Enano) GenerarArmadura() (armadura.Armadura, error) {
	return armadura.EnanoArmadura{
		Generica: armadura.Generica{
			Nombre:     "Cota de malla",
			Tipo:       "Pesada",
			Proteccion: 12,
		},
	}, nil
}

func (e Enano) GenerarCasco() (casco.Casco, error) {
	return casco.EnanoCasco{
		Generica: casco.Generica{
			Nombre:    "Yelmo",
			Habilidad: "Fiebre del Oro",
		},
	}, nil
}
