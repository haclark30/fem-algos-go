package graphs

import "slices"

type WeightedAdjacencyMatrix [][]int

func bfs(graph WeightedAdjacencyMatrix, source int, needle int) []int {
	// seen array, which nodes have i seen?
	seen := make([]bool, len(graph))
	for i := range seen {
		seen[i] = false
	}

	// prev array, where did i come from?
	prev := make([]int, len(graph))
	for i := range prev {
		prev[i] = -1
	}

	seen[source] = true
	q := []int{source}

	var curr int
	for len(q) > 0 {
		curr, q = q[0], q[1:]
		if curr == needle {
			break
		}

		adjs := graph[curr]
		for i := range adjs {
			if adjs[i] == 0 {
				continue
			}

			if seen[i] {
				continue
			}

			seen[i] = true
			prev[i] = curr
			q = append(q, i)
		}
		seen[curr] = true
	}

	// build prev back
	curr = needle
	out := []int{}

	for prev[curr] != -1 {
		out = append(out, curr)
		curr = prev[curr]
	}

	if len(out) > 0 {
		// out is from needle to source, so need to reverse it
		slices.Reverse[[]int](out)
		// source won't be in there, so add it back
		return append([]int{source}, out...)
	}
	return []int{}
}
