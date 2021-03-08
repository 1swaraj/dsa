package main

func (t *tree) addEdge(from,to int){
	tr := *t
	tr[from] = append(tr[from],Vertex{
		target: to,
	})
	tr[to] = append(tr[to],Vertex{
		target: from,
	})
}