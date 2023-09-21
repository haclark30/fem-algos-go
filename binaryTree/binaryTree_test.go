package binarytree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPreOrder(t *testing.T) {
	head := BinaryNode[int]{value: 5}
	head.left = &BinaryNode[int]{value: 3}
	head.right = &BinaryNode[int]{value: 7}

	path := PreOrder[int](&head)
	assert.Equal(t, []int{5, 3, 7}, path)
}

func TestInOrder(t *testing.T) {
	head := BinaryNode[int]{value: 5}
	head.left = &BinaryNode[int]{value: 3}
	head.right = &BinaryNode[int]{value: 7}

	path := InOrder[int](&head)
	assert.Equal(t, []int{3, 5, 7}, path)
}

func TestPostOrder(t *testing.T) {
	head := BinaryNode[int]{value: 5}
	head.left = &BinaryNode[int]{value: 3}
	head.right = &BinaryNode[int]{value: 7}

	path := PostOrder[int](&head)
	assert.Equal(t, []int{3, 7, 5}, path)
}
