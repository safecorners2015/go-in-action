package main

import (
	"fmt"
	"github.com/safecorners2015/go-in-action/chapter5/listing74/entities"
)

func main() {
	a := entities.Admin{
		Rights: 10,
	}

	a.Name = "Bill"
	a.Email = "bill@email.com"

	fmt.Printf("User: %v\n", a)
}
