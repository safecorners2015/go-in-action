package main

import (
	"fmt"

	"github.com/safecorners2015/go-in-action/chapter5/listing64/counters"
)

func main() {
	counter := counters.alertCounter(10)

	fmt.Printf("Counter: %s\n", counter)
}
