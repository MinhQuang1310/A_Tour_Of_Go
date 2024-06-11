package main

import "fmt"

type I interface {
}

func main() {

	var i I = "hello"

	s := i.(string)
	fmt.Println(s)

	s, oke := i.(string)
	fmt.Println(s, oke)

	f, oke := i.(float64)
	fmt.Println(f, oke)

	// f = i.(float64) // panic
	// fmt.Println(f)

}
