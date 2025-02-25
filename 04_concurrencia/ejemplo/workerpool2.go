package ejemplo

import (
	"fmt"
	"sync"
	"time"
)

func WorkerpoolSinBuffer() {
	tasks := []int{2, 3, 4, 5, 7, 10, 12, 40}
	numOfWorkers := 3
	var wg sync.WaitGroup

	jobs := make(chan int)
	results := make(chan int)

	for w := 1; w <= numOfWorkers; w++ {
		go worker(w, jobs, results)
	}

	go func() {
		for r := range results {
			wg.Done()
			fmt.Println("Resultado:", r)
		}
	}()

	for _, task := range tasks {
		jobs <- task
		wg.Add(1)
	}
	close(jobs)

	wg.Wait()
	close(results)
}

// simula una tarea pesada,  duerme 4 segundos y devuelve el cuadrado del numero
func worker2(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("Worker %d started job %d\n", id, job)
		cuadradoDelNumero := job * job
		time.Sleep(4 * time.Second)
		results <- cuadradoDelNumero
		fmt.Printf("Worker %d finished job %d\n", id, job)
	}

	fmt.Printf("Worker %d cerrado\n", id)
}
