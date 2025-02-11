package main

import (
	"fmt"
	"singleton/singleton"
	"sync"
)

func main() {

	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			instancia := singleton.GetInstancia(i)
			fmt.Println(i, "-", instancia.VerIntancia())
		}()
	}

	wg.Wait()

	// Las diez veces devolvió la misma instancia
}
