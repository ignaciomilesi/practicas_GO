package pieza

import "math"

type Redonda struct {
	radio int
}

func (r *Redonda) SetRadio(NuevoRadio int) {
	r.radio = NuevoRadio
}

func (r Redonda) GetRadio() int {
	return r.radio
}

type Cuadrada struct {
	lado int
}

func (c *Cuadrada) SetLado(NuevoLado int) {
	c.lado = NuevoLado
}

func (c *Cuadrada) GetLado() int {
	return c.lado
}

func AdaptadorRedondoCuadrado(c Cuadrada) Redonda {

	return Redonda{
		radio: int(math.Sqrt(float64(c.lado)) / 2),
	}

}
