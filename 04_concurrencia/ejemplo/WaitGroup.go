package ejemplo

import (
	"fmt"
	"sync"
)

func SinWaitGroup() {

	for i := 0; i < 10; i++ {

		go func() {

			fmt.Println(i)
		}()
	}

}

func UsoWaitGroup() {

	// creamos el contador
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {

		// agregamos 1 antes de tirar la goroutine
		wg.Add(1)

		go func() {

			defer wg.Done() //resta 1
			fmt.Println(i)
		}()
	}

	// bloqueamos hasta que el contador sea 0
	wg.Wait()
}
