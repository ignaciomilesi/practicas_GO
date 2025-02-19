package main

import (
	"fmt"
	"sync"
	"test/collatz"
)

func esPrimo(numero int) bool {
	return numero%2 == 0
}

func main() {

	pruebas := []int{10, 27, 100}
	var wg sync.WaitGroup

	for _, prueba := range pruebas {

		wg.Add(1)

		go func() {

			defer wg.Done()

			cantEtapas, serieResultante := collatz.CalcularSerie(prueba)

			fmt.Printf("\nPara %d fueron %d etapas y la serie resultante es %s\n", prueba, cantEtapas, serieResultante)

		}()
	}

	wg.Wait()

}
