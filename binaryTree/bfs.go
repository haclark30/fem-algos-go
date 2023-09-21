package binarytree

func bfs[T comparable](head *BinaryNode[int], needle int) bool {
	q := NewQueue[BinaryNode[int]]()
	q.enqueue(*head)

	for q.Length > 0 {
		curr, _ := q.deque()
		if curr.value == needle {
			return true
		}

		if curr.left != nil {
			q.enqueue(*curr.left)
		}
		if curr.right != nil {
			q.enqueue(*curr.right)
		}
	}

	return false
}
