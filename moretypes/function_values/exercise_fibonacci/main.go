package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int

// fibonacci returns a function that returns
//successive fibonacci numbers.
func fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
