package heap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHeap(t *testing.T) {
	heap := NewMinHeap()

	assert.Equal(t, 0, heap.Length)

	heap.insert(5)
	heap.insert(3)
	heap.insert(69)
	heap.insert(420)
	heap.insert(4)
	heap.insert(1)
	heap.insert(8)
	heap.insert(7)

	assert.Equal(t, heap.Length, 8)
	assert.Equal(t, heap.delete(), 1)
	assert.Equal(t, heap.delete(), 3)
	assert.Equal(t, heap.delete(), 4)
	assert.Equal(t, heap.delete(), 5)

	assert.Equal(t, heap.Length, 4)

	assert.Equal(t, heap.delete(), 7)
	assert.Equal(t, heap.delete(), 8)
	assert.Equal(t, heap.delete(), 69)
	assert.Equal(t, heap.delete(), 420)
	assert.Equal(t, heap.Length, 0)
	assert.Equal(t, heap.delete(), -1)
}
