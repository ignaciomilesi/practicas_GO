package raza

import (
	"Abstract_Factory/equipamiento/arma"
	"Abstract_Factory/equipamiento/armadura"
	"Abstract_Factory/equipamiento/casco"
)

type Elfo struct {
	Generico
}

func (e Elfo) GenerarArma() (arma.Arma, error) {
	return &arma.ElfoArma{
		Generica: arma.Generica{
			Nombre: "Arco",
			Tipo:   "Larga Distancia",
			Daño:   10,
		},
	}, nil
}

func (e Elfo) GenerarArmadura() (armadura.Armadura, error) {
	return armadura.ElfoArmadura{
		Generica: armadura.Generica{
			Nombre:     "Chaleco de cuero",
			Tipo:       "liviana",
			Proteccion: 5,
		},
	}, nil
}

func (e Elfo) GenerarCasco() (casco.Casco, error) {
	return casco.ElfoCasco{
		Generica: casco.Generica{
			Nombre:    "Capucha de seda",
			Habilidad: "Sigiloso",
		},
	}, nil
}
