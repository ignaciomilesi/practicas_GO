package ejemplo

import (
	"fmt"
	"sync"
	"time"
)

func SemaforoConGoroutine() {
	var wg sync.WaitGroup
	ch := make(chan int, 3)

	for i := 0; i < 40; i++ {
		wg.Add(1)
		ch <- 1
		go doSomething(i, &wg, ch)
	}

	wg.Wait()
}

func doSomething(id int, wg *sync.WaitGroup, ch chan int) {
	defer wg.Done()
	fmt.Printf("Inicio del proceso: %d\n", id)
	time.Sleep(4 * time.Second)
	fmt.Printf("Fin del proceso: %d\n", id)
	<-ch
}
