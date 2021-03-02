package main

import "fmt"

func (g *Graph) Dijkstra(n int) {
	graph := *g
	start := 8
	visited := make(map[int]int, n)
	visited[start] = 0
	for len(visited) < n {
		for k, v := range visited {
			for _, val := range graph[k] {
				_, ok := visited[val.target]
				if !ok {
					visited[val.target] = max
				}
				if v+val.weight < visited[val.target] {
					visited[val.target] = v + val.weight
				}
			}
		}
	}
	fmt.Println(visited)
}
