package main

import (
	"fmt"
	"math"
)

// tính lũy thừa x^n gán cho v
// kiểm tra v < lim thì trả về v, ngược lại trả về lim

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim

}

func main() {
	fmt.Println(pow(3, 2, 10), pow(3, 3, 20))
}
