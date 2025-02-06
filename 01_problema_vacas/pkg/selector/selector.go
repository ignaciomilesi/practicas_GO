package selector

import (
	"fmt"
	"problema_vacas/pkg/listadoVaca"
)

type Selector struct {
	ListadoVacaEnMercado listadoVaca.Lista
	LimiteCamion         int
	listaSeleccion       []int
}

// selecciona las vacas y devuelve el msj del desarrollo
func (s *Selector) SelecionarVacas(mostrarPasos bool) string {

	msj := fmt.Sprint("")
	// genera el listado de Seleccion
	s.seleccionInicial()

	if mostrarPasos {
		msj += fmt.Sprintln("------------")
	}

	// se ajusta el listado hasta que ya no se pueda mas
	for s.ajustarListadoVacas() {

		if mostrarPasos {
			msj += fmt.Sprintln("Vacas elegidas temporalmente:", s.listadosDeVacasSelecionadas())
		}
	}
	msj += fmt.Sprintln("------------")
	msj += fmt.Sprintln("Vacas seleccionadas:", s.listadosDeVacasSelecionadas())
	msj += fmt.Sprintln("Producción de vacas seleccionadas:", s.produccionTotalVacasSeleccioandas())
	msj += fmt.Sprintln("Peso de vacas seleccionadas:", s.pesoTotalVacasSeleccioandas())

	return msj
}

// seleccionara las vacas de menor producción que no se pasen del limite,
// las separa en listaSeleccionados y listaNoSeleccionados
func (s *Selector) seleccionInicial() {

	s.inicializarListaSeleccion()

	s.ListadoVacaEnMercado.OrdenarPorProduccion()

	var pesoAcumulado int

	for i, vaca := range s.ListadoVacaEnMercado.Lista {

		if pesoAcumulado+vaca.Peso <= s.LimiteCamion {

			s.listaSeleccion[i] = 1
			pesoAcumulado += vaca.Peso

		}
	}

}

// inicializa la lista que indicara si se selecciona o no la vaca de ListadoVacaEnCamion.
// 1 para seleccionado y 0 para no seleccionado. Inicializa todo en cero
func (s *Selector) inicializarListaSeleccion() {

	s.listaSeleccion = make([]int, len(s.ListadoVacaEnMercado.Lista))
}

// descarta la de mayor peso y prueba seleccionando la vaca de mayor producción no elegida
// con producción mayor a la vaca descartada. Si la encuentra actualiza el listado, sino
// prueba con la vaca que le sigue en producción. Devuelve si se pudo ajustar o no
func (s *Selector) ajustarListadoVacas() bool {

	var sePuedoAjustar bool

	//descarto la vaca
	indiceVacaDescartada := s.descartarVaca()
	vacaDescartada := s.ListadoVacaEnMercado.Lista[indiceVacaDescartada]

	pesoRestoListado := s.pesoTotalVacasSeleccioandas() - vacaDescartada.Peso

	// voy tomando las vacas de mayor producción del listado no seleccionado, ya se
	// encuentran ordenado, por lo que la ultima vaca es la de mayor producción y asi
	// para atrás
	for i := len(s.ListadoVacaEnMercado.Lista) - 1; i >= 0; i-- {

		vacaDePrueba := s.ListadoVacaEnMercado.Lista[i]

		if i == indiceVacaDescartada || // la vaca de prueba es la misma que la descartada
			s.listaSeleccion[i] == 1 || // la vaca esta seleccionada
			vacaDePrueba.Produccion < vacaDescartada.Produccion || // no genera un aumento de producción
			vacaDePrueba.Peso+pesoRestoListado > s.LimiteCamion { // supera el peso limite
			continue
		}

		// encontré una vaca que mejora la producción

		s.listaSeleccion[i] = 1 // la selecciono

		s.listaSeleccion[indiceVacaDescartada] = 0 // la descartada la deselecciono

		sePuedoAjustar = true // aviso que pude actualizar

		break
	}

	return sePuedoAjustar
}

// Devuelve el indice de la vaca que se descarta. La vaca descartada sera, de las seleccionada, la de mayor peso
// pero que su producción sea menor a la de mayor producción de la vacas no seleccionadas
func (s Selector) descartarVaca() int {

	// Primero buscamos la vaca no seleccionada de mayor producción, al estar ordenadas por producción
	// la primera que aparezca sera la que corresponda

	var mayorProduccionVacaNoSeleccionada int

	for i := len(s.listaSeleccion) - 1; i >= 0; i-- {

		vacaDePrueba := s.ListadoVacaEnMercado.Lista[i]

		if s.listaSeleccion[i] == 0 {
			mayorProduccionVacaNoSeleccionada = vacaDePrueba.Produccion
			break
		}
	}

	// Ahora buscamos la de mayor peso pero que su producción sea menor a la de vaca no seleccionada
	var mayorPeso, indiceMayorPeso int

	for i, vacaDePrueba := range s.ListadoVacaEnMercado.Lista {

		if s.listaSeleccion[i] == 1 && // esta seleccionada
			vacaDePrueba.Peso > mayorPeso && // posee mayor peso
			vacaDePrueba.Produccion < mayorProduccionVacaNoSeleccionada { // su producción en menor

			mayorPeso = vacaDePrueba.Peso
			indiceMayorPeso = i
		}
	}

	return indiceMayorPeso

}

func (s Selector) listadosDeVacasSelecionadas() []listadoVaca.Vaca {

	var listadoDeVacaSeleccioandas []listadoVaca.Vaca

	for i, selecionado := range s.listaSeleccion {

		if selecionado == 1 {

			listadoDeVacaSeleccioandas = append(listadoDeVacaSeleccioandas, s.ListadoVacaEnMercado.Lista[i])
		}
	}

	return listadoDeVacaSeleccioandas
}

func (s Selector) produccionTotalVacasSeleccioandas() int {

	var produccionTotal int

	for i, selecionado := range s.listaSeleccion {

		if selecionado == 1 {
			produccionTotal += s.ListadoVacaEnMercado.Lista[i].Produccion
		}
	}

	return produccionTotal
}

func (s Selector) pesoTotalVacasSeleccioandas() int {

	var pesoTotal int

	for i, selecionado := range s.listaSeleccion {

		if selecionado == 1 {
			pesoTotal += s.ListadoVacaEnMercado.Lista[i].Peso
		}
	}

	return pesoTotal
}
