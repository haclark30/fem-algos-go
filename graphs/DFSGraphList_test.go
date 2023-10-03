package graphs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDfsGraphList(t *testing.T) {
	list := getAdjacencyList()
	path := dfs(list, 0, 6)
	assert.Equal(t, []int{0, 1, 4, 5, 6}, path)
}

func getAdjacencyList() WeightedAdjacencyList {
	list := make(WeightedAdjacencyList, 7)
	list[0] = []GraphEdge{
		{to: 1, weight: 3},
		{to: 2, weight: 1}}

	list[1] = []GraphEdge{
		{to: 4, weight: 1},
	}
	list[2] = []GraphEdge{
		{to: 3, weight: 7},
	}

	list[3] = []GraphEdge{}

	list[4] = []GraphEdge{
		{to: 1, weight: 1},
		{to: 3, weight: 5},
		{to: 5, weight: 2},
	}

	list[5] = []GraphEdge{
		{to: 2, weight: 18},
		{to: 6, weight: 1},
	}

	list[6] = []GraphEdge{
		{to: 3, weight: 1},
	}

	return list
}
