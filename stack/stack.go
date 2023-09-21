package main

import (
	"errors"
	"fmt"
)

type node[T any] struct {
	value T
	prev  *node[T]
}

type stack[T any] struct {
	Length int
	head   *node[T]
}

func NewStack[T any]() stack[T] {
	return stack[T]{
		Length: 0,
		head:   nil,
	}
}

func (s *stack[T]) peek() (T, error) {
	if s.head == nil {
		return *new(T), errors.New("stack is empty")
	}
	return s.head.value, nil
}

func (s *stack[T]) push(item T) {
	node := node[T]{
		value: item,
	}

	s.Length++
	if s.head == nil {
		s.head = &node
		return
	}

	node.prev = s.head
	s.head = &node

}

func (s *stack[T]) pop() (T, error) {
	if s.head == nil {
		return *new(T), errors.New("stack is empty")
	}
	s.Length--

	if s.Length == 0 {
		head := *s.head
		s.head = nil
		return head.value, nil
	}

	node := *s.head
	s.head = s.head.prev

	return node.value, nil
}

func main() {
	stack := NewStack[int]()
	stack.push(0)
	stack.push(1)
	stack.push(2)

	val, _ := stack.peek()
	fmt.Println(val)

	val, _ = stack.pop()
	fmt.Println(val)

	val, _ = stack.pop()
	fmt.Println(val)
	val, _ = stack.pop()
	fmt.Println(val)
	val, err := stack.pop()
	fmt.Println(val, err)

}
