package main

import "fmt"

func main() {
	n := 9
	graph := make(Graph, n)
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
	graph.String()
	fmt.Println("Prims Algorithm")
	visted := make(map[int]bool, n)
	graph.Prims(visted, n)
	st := graph.Kruskals(n)
	fmt.Println("Kruskals Algorithm")
	st.String()
	fmt.Println("Dijkstra Algorithm")
	graph.Dijkstra(n)
}
