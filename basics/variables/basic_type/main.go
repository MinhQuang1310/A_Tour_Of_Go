package main

import (
	"fmt"
	"math/cmplx"
)

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	fmt.Printf("Type of ToBe: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type of MaxInt: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type of z: %T Value: %v\n", z, z)
}
