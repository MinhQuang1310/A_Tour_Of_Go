package main

import "fmt"

type I interface{}
type person struct {
	Name string
	Age  int
}

func (p person) String() string {
	return fmt.Sprintf("%v is %v year olds", p.Name, p.Age)
}

func main() {
	a := person{"minhquang", 22}
	b := person{"luutuan", 22}
	fmt.Println(a, "\n", b)
}
