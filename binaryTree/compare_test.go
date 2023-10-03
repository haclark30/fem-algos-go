package binarytree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCompare(t *testing.T) {
	headA := BinaryNode[int]{value: 5}
	headA.left = &BinaryNode[int]{value: 3}
	headA.right = &BinaryNode[int]{value: 45}

	headB := BinaryNode[int]{value: 5}
	headB.left = &BinaryNode[int]{value: 3}
	headB.right = &BinaryNode[int]{value: 45}

	headC := BinaryNode[int]{value: 5}
	headC.left = &BinaryNode[int]{value: 3}
	headC.left = &BinaryNode[int]{value: 45}

	assert.True(t, Compare(&headA, &headA))
	assert.True(t, Compare(&headA, &headB))

	headB.right.value = 12
	assert.False(t, Compare(&headA, &headC))
	assert.False(t, Compare(&headA, &headC))
}
