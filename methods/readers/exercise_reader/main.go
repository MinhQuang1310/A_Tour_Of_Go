package main

import (
	"golang.org/x/tour/reader"
)

type MyReader struct {
	//char byte
}

// TODO: Add a Read([]byte) (int, error) method to MyReader.
func (r MyReader) Read(b []byte) (int, error) {
	for i := range b {
		//b[i] = r.char
		b[i] = 'A'
	}
	return len(b), nil
}
func main() {
	reader.Validate(MyReader{})
}

// func main() {
// 	var input string
// 	fmt.Print("Enter a character to fill the byte slice: ")
// 	fmt.Scanf("%s", &input)
// 	if len(input) != 1 {
// 		fmt.Println("Please enter exactly one character.")
// 		return
// 	}

// 	readers := MyReader{char: input[0]}
// 	fmt.Printf("Read value: %c\n", readers.char)
// 	reader.Validate(MyReader{char: input[0]})
// }
