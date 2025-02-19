package collatz

import (
	"testing"
)

func Test_esPar(t *testing.T) {

	casos := []struct {
		entrada int
		salida  bool
	}{
		{entrada: 2, salida: true},
		{entrada: 4, salida: true},
	}

	for _, caso := range casos {
		if resultdado := esPar(caso.entrada); resultdado != caso.salida {
			t.Errorf("Error: para %d, se esperaba \"%t\" pero se obtuvo \"%t\"", caso.entrada, caso.salida, resultdado)
		}
	}
}

func Test_operaciónCollatz(t *testing.T) {

	casos := []struct {
		entrada int
		salida  int
	}{
		{entrada: 2, salida: 1},
		{entrada: 8, salida: 4},
	}

	for _, caso := range casos {
		if resultdado := operaciónCollatz(caso.entrada); resultdado != caso.salida {
			t.Errorf("Error: para %d, se esperaba %d pero se obtuvo %d", caso.entrada, caso.salida, resultdado)
		}
	}
}
