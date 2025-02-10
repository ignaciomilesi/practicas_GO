package casco

import "fmt"

type Casco interface {
	Describir()
}

type Generica struct {
	Nombre    string
	Habilidad string
}

func (g Generica) Describir() {
	fmt.Printf("Tienes un %s y posee la habilidad: %s\n", g.Nombre, g.Habilidad)
}
