package main

import "fmt"

type I interface {
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func main() {
	var i I
	describe(i)

	i = 10
	describe(i)

	i = "hello"
	describe(i)
}
