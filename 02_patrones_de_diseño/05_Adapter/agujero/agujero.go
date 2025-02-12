package agujero

import "adapter/pieza"

type Redondo struct {
	radio int
}

func (r *Redondo) SetRadio(NuevoRadio int) {
	r.radio = NuevoRadio
}

func (r Redondo) ComprobarEncaje(piezaRedonda pieza.Redonda) bool {
	return r.radio >= piezaRedonda.GetRadio()
}
