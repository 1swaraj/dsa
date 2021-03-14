package main

import "fmt"

type Edge struct {
	From int
	To int
}

type Tree map[int][]Edge

func (tr *Tree) InsertEdge(from,to int) {
	t := *tr
	t[from] = append(t[from],Edge{
		From: from,
		To:   to,
	})
	if _,ok := t[to]; !ok {
		t[to] = nil
	}
}

func (tr *Tree) Len() int {
	t := *tr
	return len(t)
}

func (tr *Tree) BFS(start int) {
	t := *tr
	visited := make(map[int]bool,t.Len())
	var queue []int
	visited[start] = true
	queue = append(queue,start)
	n:= len(queue)
	for n!=0 {
		dequeue := queue[n-1]
		fmt.Printf("%d -> ",dequeue)
		queue = queue[:n-1]
		for _,val := range t[dequeue] {
			if _,ok := visited[val.To]; !ok {
				queue = append([]int{val.To},queue...)
				visited[val.To]=true
			}
		}
		n= len(queue)
	}
}

func (tr *Tree) DFS(visited map[int]bool,start int) {
	if visited[start] == true {
		return
	}
	t := *tr
	visited[start]=true
	fmt.Printf("%d -> ",start)
	for _,val := range t[start] {
		t.DFS(visited,val.To)
	}
}

func main() {
	tree := make(Tree)
	tree.InsertEdge(0,1)
	tree.InsertEdge(0,2)
	tree.InsertEdge(1,2)
	tree.InsertEdge(2,0)
	tree.InsertEdge(3,4)
	tree.InsertEdge(2,3)
	fmt.Println(tree)
	start := 2
	fmt.Println("BFS")
	tree.BFS(start)
	fmt.Println("\nDFS")
	visited := make(map[int]bool)
	tree.DFS(visited,start)
}
