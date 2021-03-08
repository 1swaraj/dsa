package main

import "fmt"

func dfs(t map[int][]Vertex, vertex Vertex) {
	if visited[vertex.target] {
		return
	}
	visited[vertex.target] = true
	fmt.Printf("%d -> ",vertex.target)
	for _,val := range t[vertex.target] {
		dfs(t,val)
	}

}