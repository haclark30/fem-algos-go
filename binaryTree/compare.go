package binarytree

func Compare(a *BinaryNode[int], b *BinaryNode[int]) bool {
	if a == nil && b == nil {
		return true
	}

	if a == nil || b == nil {
		return false
	}

	if a.value != b.value {
		return false
	}

	return Compare(a.left, b.left) && Compare(a.right, b.right)

}
