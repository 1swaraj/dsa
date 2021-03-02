package main

import "fmt"

func (g *Graph) AddEdge(u, v, w int) {
	graph := *g
	graph[u] = append(graph[u], Edge{
		target: v,
		weight: w,
		from:   u,
	})
	graph[v] = append(graph[v], Edge{
		target: u,
		weight: w,
		from:   v,
	})
}

func (g *Graph) String() {
	graph := *g
	for key, val := range graph {
		fmt.Printf("%d %v\n", key, val)
	}
}
