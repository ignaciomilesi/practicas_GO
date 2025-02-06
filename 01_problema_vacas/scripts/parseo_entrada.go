package scripts

import (
	"strconv"
	"strings"
)

func ParseoEntrada(entrada string) (limitPesoCamion int, listaPeso, listaProduccion []int) {

	entradaSinEspacios := strings.Fields(entrada)

	// la entrada es [cant limitPesoCamion listaPesoVaca listaProducVaca]

	if i, err := strconv.Atoi(entradaSinEspacios[1]); err == nil {
		limitPesoCamion = i
	}

	listaPesoString := strings.Split(entradaSinEspacios[2], ",")
	listaProduccionString := strings.Split(entradaSinEspacios[3], ",")

	cantVacas, _ := strconv.Atoi(entradaSinEspacios[0])

	for i := 0; i < cantVacas; i++ {

		if peso, err := strconv.Atoi(listaPesoString[i]); err == nil {
			listaPeso = append(listaPeso, peso)
		}

		if produccion, err := strconv.Atoi(listaProduccionString[i]); err == nil {
			listaProduccion = append(listaProduccion, produccion)
		}

	}

	return limitPesoCamion, listaPeso, listaProduccion
}
