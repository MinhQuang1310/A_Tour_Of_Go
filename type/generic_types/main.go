package main

import "fmt"

// List represents a singly-linked list that holds
// values of any type.
type List[T any] struct {
	next *List[T]
	val  T
}

// New constructs a new List with the given value.
func New[T any](val T) *List[T] {
	return &List[T]{val: val}
}

// Append appends a new node to the end of the list
// with the given value.
func (l *List[T]) Append(val T) *List[T] {
	if l == nil {
		return New(val)
	}
	for ; l.next != nil; l = l.next {
	}
	l.next = New(val)
	return l
}

// String returns a string representation of the list.
func (l *List[T]) String() string {
	s := ""
	for ; l != nil; l = l.next {
		s += fmt.Sprintf("%v -> ", l.val)
	}
	return s + "nil"
}

func main() {
	l := New(1)
	l.Append(2).Append(3)
	fmt.Println(l)
}
