package maze

type Point struct {
	x int
	y int
}

var dir = []Point{
	{-1, 0},
	{1, 0},
	{0, -1},
	{0, 1},
}

func walk(maze []string, wall string, curr Point, end Point, seen [][]bool, path []Point) ([]Point, bool) {
	// base case off map
	if (curr.x < 0 || curr.x > len(maze[0])) ||
		(curr.y < 0 || curr.y > len(maze)) {
		return path, false
	}
	// on wall

	if maze[curr.y][curr.x] == wall[0] {
		return path, false
	}

	// at end
	if curr.x == end.x && curr.y == end.y {
		path = append(path, curr)
		return path, true
	}

	// seen
	if seen[curr.y][curr.x] {
		return path, false
	}

	// pre
	seen[curr.y][curr.x] = true
	path = append(path, curr)

	// recurse
	for _, currDir := range dir {
		path, done := walk(maze, wall, Point{curr.x + currDir.x, curr.y + currDir.y}, end, seen, path)
		if done {
			return path, true
		}
	}
	// post
	path = path[:len(path)-1]

	return path, false
}

func Solve(maze []string, wall string, start Point, end Point) []Point {
	seen := make([][]bool, len(maze))
	var path []Point

	for i := range seen {
		seen[i] = make([]bool, len(maze[0]))
	}

	path, _ = walk(maze, wall, start, end, seen, path)

	return path

}
