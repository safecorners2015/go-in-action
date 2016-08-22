package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	// Create a Buffer value and write a string to the buffer.
	// Using the Write method that implements io.Writer.
	var b bytes.Buffer
	b.Write([]byte("Hello "))

	// Use Fprintf to concatenate a string to the Buffer.
	// Passing the address of a bytes.Buffer value for io.Writer.
	fmt.Fprintf(&b, "World!")

	b.WriteTo(os.Stdout)
}
