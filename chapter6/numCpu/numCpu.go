package main

import (
	"fmt"
	"runtime"
)

func main() {
	// Set logical processor on each available core
	// runtime.GOMAXPROCS(runtime.NumCPU())

	fmt.Println(runtime.NumCPU())

}
