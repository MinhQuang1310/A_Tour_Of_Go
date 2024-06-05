package main

import (
	"fmt"
	"math/rand"
)

// main is the entry point of the program.
func generateRandomNumber() int {
	return rand.Intn(10)
}

func main() {
	randomNumber := generateRandomNumber()
	fmt.Println("Random number:", randomNumber)
}
