package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	r := strings.NewReader("Hello, Reader!")
	b := make([]byte, 8)
	for {
		n, error := r.Read(b)
		fmt.Printf("n = %v b = %v error = %v\n", n, b, error)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if error == io.EOF {
			break
		}
	}
}
