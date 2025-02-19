package collatz

import "strconv"

func esPar(numero int) bool {
	return numero%2 == 0
}

func operaciónCollatz(numero int) int {

	if esPar(numero) {
		return numero / 2

	} else {
		return 3*numero + 1
	}
}

func CalcularSerie(entrada int) (cantEtapas int, serieResultante string) {

	serieResultante += strconv.Itoa(entrada)

	for entrada != 1 {
		entrada = operaciónCollatz(entrada)
		serieResultante += ", " + strconv.Itoa(entrada)
		cantEtapas += 1
	}

	return
}
