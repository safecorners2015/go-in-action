// simple API Server for testing endpoint
package main

import (
	"log"
	"net/http"

	"github.com/safecorners2015/go-in-action/chapter9/listing17/handlers"
)

// entry point
func main() {
	handlers.Routes()

	log.Println("listner: Started : Listening on: 4000")
	http.ListenAndServe(":4000", nil)
}
