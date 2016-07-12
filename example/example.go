package main

import (
	"fmt"
	"github.com/bramvdbogaerde/go-uniquestring"
)

func main() {
	// Generate a random unique string
	s := uniquestring.New()
	fmt.Println(s)
}
