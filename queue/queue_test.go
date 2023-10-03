package queue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueue(t *testing.T) {
	list := NewQueue[int]()
	list.enqueue(5)
	list.enqueue(7)
	list.enqueue(9)

	val, _ := list.deque()
	assert.Equal(t, 5, val)
	assert.Equal(t, 2, list.Length)

	list.enqueue(11)
	val, _ = list.deque()
	assert.Equal(t, 7, val)
	val, _ = list.deque()
	assert.Equal(t, 9, val)
	val, _ = list.peek()
	assert.Equal(t, 11, val)
	val, _ = list.deque()
	assert.Equal(t, 11, val)
	val, err := list.deque()
	assert.NotNil(t, err)
	assert.Equal(t, 0, list.Length)

	list.enqueue(69)
	val, err = list.peek()
	assert.Equal(t, 69, val)
	assert.Equal(t, 1, list.Length)
}
