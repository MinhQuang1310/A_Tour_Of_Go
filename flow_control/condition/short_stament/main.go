package main

import (
	"fmt"
	"math"
)

func show_stament(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}
func main() {
	fmt.Println(show_stament(3, 2, 10), show_stament(3, 3, 20))
}
