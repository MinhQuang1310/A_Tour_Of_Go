package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}
	Walk(t.Left, ch)
	ch <- t.Value
	Walk(t.Right, ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go func() {
		Walk(t1, ch1)
		close(ch1)
	}()
	go func() {
		Walk(t2, ch2)
		close(ch2)
	}()
	for v1 := range ch1 {
		v2, ok := <-ch2
		if !ok || v1 != v2 {
			return false
		}
	}
	_, ok := <-ch2
	return !ok
}

func main() {
	ch := make(chan int)
	go func() {
		Walk(tree.New(1), ch)
		close(ch)
	}()
	for v := range ch {
		fmt.Println(v)
	}

	// Example usage of Same function
	t1 := tree.New(1)
	t2 := tree.New(1)
	fmt.Println("t1 and t2 have the same values:", Same(t1, t2))
}
