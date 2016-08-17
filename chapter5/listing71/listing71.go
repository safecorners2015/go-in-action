package main

import (
	"fmt"
	"github.com/safecorners2015/go-in-action/chapter5/listing71/entities"
)

func main() {
	u := entities.User{
		Name:  "Bill",
		email: "bill@example.com",
	}

	fmt.Printf("User: %v\n", u)
}
