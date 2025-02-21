package ejemplo

import (
	"fmt"
	"time"
)

func CreandoChannel() {
	c := make(chan int)

	go func(ch chan int, numero int) {

		ch <- numero
		fmt.Println("Numero Colocado")

	}(c, 10)

	fmt.Print(<-c)
}

func BloqueoChannelConSleep() {
	c := make(chan int)

	go func(ch chan int, numero int) {

		ch <- numero
		fmt.Println("Numero Colocado")

	}(c, 10)

	time.Sleep(3 * time.Second)
	fmt.Print("hola")

}

func BloqueoChannelConOtroChannel() {
	dato := make(chan int)
	finTarea := make(chan bool)

	go func(c_dato chan int, c_finTarea chan bool, numero int) {

		c_dato <- numero
		fmt.Println("Numero Colocado")
		c_finTarea <- true

	}(dato, finTarea, 10)

	<-finTarea
	fmt.Print(<-dato)
}

func DeadlockPorDobleenvio() {
	c := make(chan int)

	go func(ch chan int) {

		fmt.Println(<-ch)

	}(c)

	c <- 10
	c <- 10

	time.Sleep(100 * time.Millisecond) //solo para darle tiempo a que se ejecute la goroutine
}
