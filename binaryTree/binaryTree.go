package binarytree

type BinaryNode[T comparable] struct {
	value T
	left  *BinaryNode[T]
	right *BinaryNode[T]
}

func walkPreOrder[T comparable](curr *BinaryNode[T], path *[]T) []T {
	if curr == nil {
		return *path
	}

	*path = append(*path, curr.value)
	walkPreOrder[T](curr.left, path)
	walkPreOrder[T](curr.right, path)

	return *path
}

func PreOrder[T comparable](head *BinaryNode[T]) []T {
	path := make([]T, 0)
	return walkPreOrder[T](head, &path)
}

func walkInOrder[T comparable](curr *BinaryNode[T], path *[]T) []T {
	if curr == nil {
		return *path
	}

	walkInOrder[T](curr.left, path)
	*path = append(*path, curr.value)
	walkInOrder[T](curr.right, path)

	return *path
}

func InOrder[T comparable](head *BinaryNode[T]) []T {
	path := make([]T, 0)
	return walkInOrder[T](head, &path)
}

func walkPostOrder[T comparable](curr *BinaryNode[T], path *[]T) []T {
	if curr == nil {
		return *path
	}

	walkPostOrder[T](curr.left, path)
	walkPostOrder[T](curr.right, path)
	*path = append(*path, curr.value)

	return *path
}

func PostOrder[T comparable](head *BinaryNode[T]) []T {
	path := make([]T, 0)
	return walkPostOrder[T](head, &path)
}
