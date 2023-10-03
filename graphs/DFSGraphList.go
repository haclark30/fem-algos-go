package graphs

type GraphEdge struct {
	to     int
	weight int
}

type WeightedAdjacencyList [][]GraphEdge

// TODO: change this to return slice instead of use slice pointer
func walk(graph WeightedAdjacencyList, curr int, needle int, seen []bool, path *[]int) bool {

	if seen[curr] {
		return false
	}

	seen[curr] = true

	//pre
	*path = append(*path, curr)
	if curr == needle {
		return true
	}
	// recurse
	list := graph[curr]
	for i := range list {
		edge := list[i]

		if walk(graph, edge.to, needle, seen, path) {
			return true
		}
	}
	//post
	*path = (*path)[:len(*path)-1]
	return false

}
func dfs(graph WeightedAdjacencyList, source int, needle int) []int {
	seen := make([]bool, len(graph))
	path := make([]int, 0)

	walk(graph, source, needle, seen, &path)
	if len(path) > 0 {
		return path
	}
	return nil
}
