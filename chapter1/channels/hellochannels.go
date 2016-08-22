package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func printer(ch chan int) {
	for i := range ch {
		fmt.Println("Received ", i)
	}
	wg.Done()
}

// entry point
func main() {
	c := make(chan int)
	go printer(c)
	wg.Add(1)

	// Send 10 integers on the channel.
	for i := 1; i <= 10; i++ {
		c <- i
	}

	close(c)
	wg.Wait()
}
