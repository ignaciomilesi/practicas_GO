package ejemplo

import (
	"fmt"
	"time"
)

func Workerpool() {
	tasks := []int{2, 3, 4, 5, 7, 10, 12, 40}
	numOfWorkers := 3

	jobs := make(chan int, len(tasks))
	results := make(chan int, len(tasks))

	for w := 1; w <= numOfWorkers; w++ {
		go worker(w, jobs, results)
	}

	for _, task := range tasks {
		jobs <- task
	}
	close(jobs)

	for i := 0; i < len(tasks); i++ {
		<-results
	}

	// solo para dar tiempo a cerrar los worker
	time.Sleep(2 * time.Second)
}

// simula una tarea pesada,  duerme 4 segundos y devuelve el cuadrado del numero
func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("Worker %d started job %d\n", id, job)
		cuadradoDelNumero := job * job
		time.Sleep(4 * time.Second)
		results <- cuadradoDelNumero
		fmt.Printf("Worker %d finished job %d\n", id, job)
	}

	fmt.Printf("Worker %d cerrado\n", id)
}
