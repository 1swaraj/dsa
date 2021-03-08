package main

import "fmt"

func bfs(t map[int][]Vertex,vertex Vertex,n int) {
	var queue []int
	visited := make(map[int]bool,n)
	visited[vertex.target]=true
	queue = append(queue,vertex.target)
	for len(queue)>0 {
		length := len(queue)
		dequeue := queue[length-1]
		queue = queue[:length-1]
		fmt.Printf("%d -> ",dequeue)
		for _,val := range t[dequeue] {
			if visited[val.target] {
				continue
			}
			visited[val.target]=true
			queue = append([]int{val.target},queue...)
		}
	}
}