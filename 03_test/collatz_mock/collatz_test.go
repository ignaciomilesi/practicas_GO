package collatz_mock

import (
	"testing"
)

func Test_operaciónCollatz(t *testing.T) {

	casos := []struct {
		entrada  int
		salida   int
		mockFunc func()
	}{
		{entrada: 2, salida: 1, mockFunc: func() {
			esPar = func(int) bool { return false }

			// si hubiera otra función la cargo aca esPar2 = func(){....}
		},
		},
		{entrada: 8, salida: 4, mockFunc: func() {
			esPar = func(int) bool { return false }
		},
		},
	}

	originalEsPar := esPar // me guardo la función original

	for _, caso := range casos {

		caso.mockFunc() // aca realizo el mockeo de la función

		if resultdado := operaciónCollatz(caso.entrada); resultdado != caso.salida {
			t.Errorf("Error: para %d, se esperaba %d pero se obtuvo %d", caso.entrada, caso.salida, resultdado)
		}
	}

	esPar = originalEsPar //le saco el mock, para que el siguiente test no tenga la función mockeada
}
