package arma

import "fmt"

type Arma interface {
	Usar()
	Describir()
	Mejorar()
}

type Generica struct {
	Nombre string
	Tipo   string // Corto o largo alcance
	Daño   int
}

func (g Generica) Usar() {
	fmt.Printf("Atacas con %s. Haces %d de daño \n", g.Nombre, g.Daño)
}

func (g Generica) Describir() {
	fmt.Printf("Tienes un %s. Es un arma del Tipo %s\n", g.Nombre, g.Tipo)
}

func (g *Generica) Mejorar() {
	g.Daño += 3
	fmt.Printf("Mejoras tu %s. Tu nuevo daño es %d\n", g.Nombre, g.Daño)
}
