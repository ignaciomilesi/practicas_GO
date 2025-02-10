package armadura

import "fmt"

type Armadura interface {
	Describir()
}

type Generica struct {
	Nombre     string
	Tipo       string // Liviana o pesada
	Proteccion int
}

func (g Generica) Describir() {
	fmt.Printf("Tienes un %s, es una armadura tipo %s y posee %d de protección\n", g.Nombre, g.Tipo, g.Proteccion)
}
