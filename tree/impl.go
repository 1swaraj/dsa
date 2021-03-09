package main

import "fmt"

func main() {
	t := make(tree)
	visited = make(map[int]bool)
	t.addEdge(1, 8)
	t.addEdge(8, 5)
	t.addEdge(1, 7)
	t.addEdge(5, 10)
	t.addEdge(5, 6)
	t.addEdge(10, 6)
	dfs(t, Vertex{
		target: 0,
	})
	fmt.Println()
	bfs(t, Vertex{
		target: 0,
	}, 6)
}
