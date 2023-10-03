package graphs

import (
	"math"
	"slices"
)

func hasUnvisited(seen []bool, dists []float64) bool {
	for i, s := range seen {
		if !s && dists[i] < math.Inf(1) {
			return true
		}
	}
	return false
}

func getLowestUnvisited(seen []bool, dists []float64) int {
	idx := -1
	lowestDist := math.Inf(1)

	for i := range seen {
		if seen[i] {
			continue
		}

		if lowestDist > dists[i] {
			lowestDist = dists[i]
			idx = i
		}
	}
	return idx
}
func DijkstraList(source int, sink int, arr WeightedAdjacencyList) []int {
	seen := make([]bool, len(arr))
	prev := make([]int, len(arr))
	dists := make([]float64, len(arr))

	for i := range dists {
		dists[i] = math.Inf(1)
	}

	for i := range prev {
		prev[i] = -1
	}

	dists[source] = 0

	for hasUnvisited(seen, dists) {
		curr := getLowestUnvisited(seen, dists)
		seen[curr] = true

		adjs := arr[curr]
		for i := range adjs {
			edge := adjs[i]
			if seen[edge.to] {
				continue
			}

			dist := dists[curr] + float64(edge.weight)
			if dist < dists[edge.to] {
				dists[edge.to] = dist
				prev[edge.to] = curr
			}
		}
	}

	out := make([]int, 0)
	curr := sink

	for prev[curr] != -1 {
		out = append(out, curr)
		curr = prev[curr]
	}

	out = append(out, source)
	slices.Reverse(out)
	return out
}
