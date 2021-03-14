package main

import (
	"fmt"
	"math"
)

type Edge struct {
	To     int
	From   int
	Weight int
}

type Graph map[int][]Edge

func (g *Graph) Len() int {
	graph := *g
	return len(graph)
}

func (graph *Graph) AddEdge(f, t, w int) {
	g := *graph
	g[f] = append(g[f], Edge{
		To:     t,
		From:   f,
		Weight: w,
	})
	g[t] = append(g[t], Edge{
		To:     f,
		From:   t,
		Weight: w,
	})
}

func main() {
	graph := make(Graph)
	graph.AddEdge(1, 2, 8)
	graph.AddEdge(0, 1, 4)
	graph.AddEdge(2, 3, 7)
	graph.AddEdge(3, 4, 9)
	graph.AddEdge(4, 5, 10)
	graph.AddEdge(5, 6, 2)
	graph.AddEdge(6, 7, 1)
	graph.AddEdge(7, 8, 7)
	graph.AddEdge(7, 0, 8)
	graph.AddEdge(1, 7, 11)
	graph.AddEdge(2, 8, 2)
	graph.AddEdge(8, 6, 6)
	graph.AddEdge(2, 5, 4)
	graph.AddEdge(3, 5, 14)
	n := graph.Len()
	adjMat := make([][]int,n)
	for key,val := range graph {
		adjMat[key] = make([]int,n)
		for j:=0;j<n;j++{
			if key==j {
				continue
			}
			adjMat[key][j]=math.MaxInt16
		}
		for _,v:= range val {
			adjMat[key][v.To]=v.Weight
		}
	}
	for k:=0;k<n;k++ {
		for i:=0;i<n;i++ {
			for j:=0;j<n;j++{
				adjMat[i][j]= min(adjMat[i][j],adjMat[i][k]+adjMat[k][j])
			}
		}
	}
	fmt.Println(adjMat)
}

func min(x,y int) int {
	if x<y {
		return x
	}
	return y
}

