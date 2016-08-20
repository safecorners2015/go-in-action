// Unbuffered Channel
package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	baton := make(chan int)

	wg.Add(1)

	go Runner(baton)

	baton <- 1

	wg.Wait()
}

func Runner(baton chan int) {
	var newRunner int

	runner := <-baton

	fmt.Printf("%d runner taken baton & run\n", runner)

	if runner != 4 {
		newRunner = runner + 1
		fmt.Printf("%d runner is waiting\n", newRunner)
		go Runner(baton)
	}

	time.Sleep(100 * time.Millisecond)

	if runner == 4 {
		fmt.Printf("%d runner arrived. Race is finished\n", runner)
		wg.Done()
		return
	}

	fmt.Printf("%d runner give a baton  to %d runner\n", runner, newRunner)

	baton <- newRunner
}
