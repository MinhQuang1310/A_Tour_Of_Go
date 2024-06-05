package main

import (
	"fmt"
	"strings"
)

func main() {
	x := [][]string{
		{"-", "-", "-"},
		{"-", "-", "-"},
		{"-", "-", "-"},
	}

	x[0][0] = "X"
	x[0][1] = "O"
	x[0][2] = "O"
	x[1][0] = "O"
	x[1][1] = "X"
	x[1][2] = "O"
	x[2][0] = "X"
	x[2][1] = "O"
	x[2][2] = "X"
	for i := 0; i < len(x); i++ {
		fmt.Printf("%s\n", strings.Join(x[i], " "))
	}
}
