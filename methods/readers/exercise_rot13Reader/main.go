package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func rot13(b byte) byte {
	if b >= 'a' && b <= 'm' {
		b += 13
	} else if b >= 'A' && b <= 'M' {
		b += 13
	} else if b >= 'n' && b <= 'z' {
		b -= 13
	} else if b >= 'N' && b <= 'Z' {
		b -= 13
	}
	return b
}

func (r *rot13Reader) Read(b []byte) (int, error) {
	n, err := r.r.Read(b)
	for i := range b {
		b[i] = rot13(b[i])
	}
	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
