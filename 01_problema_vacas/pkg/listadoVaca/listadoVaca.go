package listadoVaca

import (
	"sort"
)

type Vaca struct {
	Numero     int
	Peso       int
	Produccion int
}

type Lista struct {
	Lista []Vaca
}

// Creara el listado de vacas dada el listado de pesos y producción
func (l *Lista) Crear(listaPeso, listaProduccion []int) {

	for i, peso := range listaPeso {
		nuevaVaca := Vaca{
			Numero:     i + 1,
			Peso:       peso,
			Produccion: listaProduccion[i],
		}
		l.Lista = append(l.Lista, nuevaVaca)
	}
}

func (l *Lista) OrdenarPorProduccion() {

	sort.Slice(l.Lista, func(i, j int) bool {
		return l.Lista[i].Produccion < l.Lista[j].Produccion
	})

}
