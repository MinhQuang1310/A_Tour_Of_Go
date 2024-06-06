package main

import (
	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
	// Create a 2D slice of uint8 values with a length of dy.
	// Each element of the slice is a slice of uint8 values with a length of dx.
	pic := make([][]uint8, dy)

	// Iterate over each element in the outer slice (rows).
	for i := 0; i < dy; i++ {
		// Create a new slice of uint8 values with a length of dx.
		// This will be the row of the 2D slice.
		pic[i] = make([]uint8, dx)

		// Iterate over each element in the inner slice (columns).
		for j := 0; j < dx; j++ {
			// Calculate the value for the current pixel by taking the sum of i and j and taking the modulus of 256.
			// This effectively creates a grayscale image where the value at each pixel is determined by the sum of its row and column indices.
			pic[i][j] = uint8((i + j) % 256)
		}
	}

	// Return the 2D slice of uint8 values.
	return pic
}

func main() {
	pic.Show(Pic)
}
