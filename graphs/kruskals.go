package main

import (
	"container/heap"
)

func (v Vertices) Len() int {
	return len(v)
}

func (v Vertices) Less(i, j int) bool {
	return v[i].weight < v[j].weight
}

func (v Vertices) Swap(i, j int) {
	v[i], v[j] = v[j], v[i]
}

func (v *Vertices) Push(x interface{}) {
	vertex := x.(Vertex)

	*v = append(*v, vertex)
}

func (v *Vertices) Pop() interface{} {
	values := *v
	length := len(values)
	res := values[length-1]
	*v = values[:length-1]
	return res
}

func (g *Graph) Kruskals(n int) Graph {
	graph := *g
	vertices := Vertices{}
	st := make(Graph, n)
	for _, v := range graph {
		for _, val := range v {
			vertex := Vertex{
				from:   val.from,
				to:     val.target,
				weight: val.weight,
			}
			vertices = append(vertices, vertex)
		}
	}
	heap.Init(&vertices)
	set := make(map[Vertex]bool)
	for n > 0 {
		pop := heap.Pop(&vertices).(Vertex)
		_, ok := set[pop]
		if ok {
			continue
		}
		set[pop] = true
		st.AddEdge(pop.from, pop.to, pop.weight)
		pop.from, pop.to = pop.to, pop.from
		set[pop] = true
		n--
	}
	return st
}
