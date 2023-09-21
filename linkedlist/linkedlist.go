package linkedlist

import "errors"

type LinkedList[T any] interface {
	prepend(item T)
	insertAt(item T, idx int) error
	append(item T)
	remove(item T) (T, error)
	get(idx int) (T, error)
	removeAt(idx int) (T, error)
}

type ListNode[T any] struct {
	value T
	prev  *ListNode[T]
	next  *ListNode[T]
}

type DoublyLinkedList[T comparable] struct {
	Length int
	head   *ListNode[T]
	tail   *ListNode[T]
}

func NewDoublyLinkedList[T comparable]() DoublyLinkedList[T] {
	return DoublyLinkedList[T]{
		Length: 0,
		head:   nil}
}

func (l *DoublyLinkedList[T]) prepend(item T) {
	node := ListNode[T]{value: item}

	l.Length++
	if l.head == nil {
		l.head = &node
		l.tail = &node
		return
	}

	node.next = l.head
	l.head.prev = &node

	l.head = &node
}

func (l *DoublyLinkedList[T]) insertAt(item T, idx int) error {
	if idx > l.Length {
		return errors.New("index greater than length of list")
	} else if idx == l.Length {
		l.append(item)
		return nil
	} else if idx == 0 {
		l.prepend(item)
		return nil
	}

	l.Length++
	curr, err := l.getAt(idx)
	if err != nil {
		return err
	}

	node := ListNode[T]{value: item}

	node.next = curr
	node.prev = curr.prev

	curr.prev.next = &node
	curr.prev = &node

	// this was prime's implementation, which is wrong
	// curr.prev = &node
	// if node.prev != nil {
	// 	node.prev.next = curr
	// }

	return nil
}

func (l *DoublyLinkedList[T]) append(item T) {
	l.Length++
	node := ListNode[T]{value: item}

	if l.tail == nil {
		l.head = &node
		l.tail = &node
	}
	node.prev = l.tail
	l.tail.next = &node
	l.tail = &node
}

func (l *DoublyLinkedList[T]) remove(item T) (T, error) {
	curr := l.head
	for i := 0; curr != nil && i < l.Length; i++ {
		if curr.value == item {
			break
		}
		curr = curr.next
	}

	if curr == nil {
		return *new(T), errors.New("value not in list")
	}
	return l.removeNode(curr)
}

func (l *DoublyLinkedList[T]) get(idx int) (T, error) {
	node, err := l.getAt(idx)
	if err != nil {
		return *new(T), err
	}
	return node.value, nil
}

func (l *DoublyLinkedList[T]) getAt(idx int) (*ListNode[T], error) {
	if idx > l.Length {
		return nil, errors.New("index greater than length")
	}

	curr := l.head
	for i := 0; curr != nil && i < idx; i++ {
		curr = curr.next
	}

	return curr, nil
}

func (l *DoublyLinkedList[T]) removeAt(idx int) (T, error) {
	node, err := l.getAt(idx)
	if err != nil {
		return *new(T), err
	}

	return l.removeNode(node)
}

func (l *DoublyLinkedList[T]) removeNode(curr *ListNode[T]) (T, error) {
	l.Length--

	if l.Length == 0 {
		l.head = nil
		l.tail = nil
		return curr.value, nil
	}

	if curr.prev != nil {
		curr.prev.next = curr.next
	}

	if curr.next != nil {
		curr.next.prev = curr.prev
	}

	if curr == l.head {
		l.head = curr.next
	}

	if curr == l.tail {
		l.tail = curr.prev
	}

	curr.prev = nil
	curr.next = nil
	return curr.value, nil
}
