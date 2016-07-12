package main

import (
	"fmt"
	"github.com/bramvdbogaerde/go-randomstring"
)

func main() {
	// Generate a random unique string
	s := randomstring.New()
	fmt.Println(s)
}
