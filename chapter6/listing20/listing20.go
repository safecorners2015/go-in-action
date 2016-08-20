// Unbuffered
// Tenis
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	// Declare Unbuffered channel
	court := make(chan int)

	wg.Add(2)

	go player("nadal", court)
	go player("djokovic", court)

	court <- 1

	wg.Wait()
}

func player(name string, court chan int) {
	defer wg.Done()

	for {

		ball, ok := <-court
		if !ok {
			// player win when channnel is closed
			fmt.Printf("%s is won!\n", name)
			return
		}

		n := rand.Intn(100)
		if n%13 == 0 {
			fmt.Printf("player %s is failed to receive\n", name)

			close(court)
			return
		}

		fmt.Printf("player %s returns [%d]ball\n", name, ball)
		ball++

		court <- ball
	}
}
