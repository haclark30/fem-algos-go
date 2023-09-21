package binarytree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBfs(t *testing.T) {
	head := BinaryNode[int]{value: 5}
	head.left = &BinaryNode[int]{value: 3}
	head.right = &BinaryNode[int]{value: 7}

	assert.Equal(t, true, bfs[BinaryNode[int]](&head, 3))
	assert.Equal(t, true, bfs[BinaryNode[int]](&head, 7))
	assert.Equal(t, false, bfs[BinaryNode[int]](&head, 15))

}
