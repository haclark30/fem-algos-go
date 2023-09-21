package main

import (
	"errors"
	"fmt"
)

type Node[T any] struct {
	value T
	next  *Node[T]
}

type Queue[T any] struct {
	Length int
	head   *Node[T]
	tail   *Node[T]
}

func NewQueue[T any]() Queue[T] {
	return Queue[T]{
		Length: 0,
		head:   nil,
		tail:   nil,
	}
}

func (q *Queue[T]) peek() (T, error) {
	if q.head != nil {
		return q.head.value, nil
	}
	return *new(T), errors.New("queue is empty")
}

func (q *Queue[T]) deque() (T, error) {
	if q.head == nil {
		return *new(T), errors.New("queue is empty")
	}

	q.Length--

	head := q.head
	q.head = q.head.next
	if q.head == nil {
		q.tail = nil
	}
	return head.value, nil
}

func (q *Queue[T]) enqueue(item T) {
	q.Length++
	node := Node[T]{value: item}
	if q.tail == nil {
		q.tail = &node
		q.head = &node
		return
	}

	q.tail.next = &node
	q.tail = &node
}

func main() {
	queue := NewQueue[int]()
	queue.enqueue(5)
	queue.enqueue(3)
	queue.enqueue(2)

	val, _ := queue.deque()
	fmt.Println(val)
	val, _ = queue.deque()
	fmt.Println(val)
	val, _ = queue.deque()
	fmt.Println(val)
	val, err := queue.deque()
	fmt.Println(err)

}
