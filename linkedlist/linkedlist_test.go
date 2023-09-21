package linkedlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestList(t *testing.T) {
	linkedList := NewDoublyLinkedList[int]()
	err := testList(t, &linkedList)
	if err != nil {
		t.Fatal(err)
	}
}

func testList(t *testing.T, list LinkedList[int]) error {
	list.prepend(5)
	list.append(3)
	list.insertAt(7, 1)

	val, err := list.get(1)
	assert.Equal(t, 7, val)
	assert.Nil(t, err, "error should be nil")

	val, err = list.remove(7)
	assert.Equal(t, 7, val)
	assert.Nil(t, err)
	return nil
}
