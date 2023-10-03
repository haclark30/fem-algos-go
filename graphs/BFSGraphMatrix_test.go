package graphs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var matrix = WeightedAdjacencyMatrix{
	{0, 3, 1, 0, 0, 0, 0}, // 0
	{0, 0, 0, 0, 1, 0, 0},
	{0, 0, 7, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0},
	{0, 1, 0, 5, 0, 2, 0},
	{0, 0, 18, 0, 0, 0, 1},
	{0, 0, 0, 1, 0, 0, 1},
}

func TestBfsMatrix(t *testing.T) {
	path := bfs(matrix, 0, 6)
	assert.Equal(t, path, []int{0, 1, 4, 5, 6})
}
