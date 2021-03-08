package main

import "fmt"

func main() {
	t := make(tree)
	visited = make(map[int]bool)
	t.addEdge(0, 1)
	t.addEdge(2, 1)
	t.addEdge(0, 3)
	t.addEdge(0, 4)
	t.addEdge(2, 5)
	dfs(t, Vertex{
		target: 0,
	})
	fmt.Println()
	bfs(t, Vertex{
		target: 0,
	}, 6)
}
