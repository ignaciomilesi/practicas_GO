package ejemplo

import (
	"fmt"
	"time"
)

func SinnMultiplex() {
	c1 := make(chan int)
	c2 := make(chan int)

	duration1 := 4 * time.Second
	duration2 := 2 * time.Second

	go doSomething2(duration1, c1, 1)
	go doSomething2(duration2, c2, 2)

	fmt.Printf("Channel 1 received: %d\n", <-c1)
	fmt.Printf("Channel 2 received: %d\n", <-c2)

}

func ConMultiplex() {
	c1 := make(chan int)
	c2 := make(chan int)

	duration1 := 4 * time.Second
	duration2 := 2 * time.Second

	go doSomething2(duration1, c1, 1)
	go doSomething2(duration2, c2, 2)

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Printf("Channel 1 received: %d\n", msg1)
		case msg2 := <-c2:
			fmt.Printf("Channel 2 received: %d\n", msg2)
		}
	}
}

func doSomething2(duration time.Duration, channel chan int, param int) {
	time.Sleep(duration)
	channel <- param
}
