package binarytree

func search(curr *BinaryNode[int], needle int) bool {
	if curr == nil {
		return false
	}

	if curr.value == needle {
		return true
	}

	if curr.value < needle {
		return search(curr.right, needle)
	}
	return search(curr.left, needle)
}

func dfs(head BinaryNode[int], needle int) bool {
	return search(&head, needle)
}
