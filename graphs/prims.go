package main

import "fmt"

func (g *Graph) Prims(visited map[int]bool, n int) {
	sum := 0
	graph := *g
	start := 3
	fmt.Printf("S -> %d\n", start)
	visited[start] = true
	for i := 1; i < n; i++ {
		min := max
		minV := -1
		vertex := -1
		for k, _ := range visited {
			for _, val := range graph[k] {
				_, ok := visited[val.target]
				if !ok && val.weight < min {
					min = val.weight
					minV = val.target
					vertex = val.from
				}
			}
		}
		visited[minV] = true
		sum += min
		fmt.Printf("%d->%d %d\n", vertex, minV, min)
	}
	fmt.Printf("\nSum :=> %d\n", sum)
}
