package main

import (
	"container/heap"
	"fmt"
)

type maxHeap []string

func (m *maxHeap) Pop() interface{} {
	values := *m
	length := len(values)
	*m = values[:length-1]
	return values[length-1]
}

func (m *maxHeap) Push(x interface{}) {
	*m = append(*m,x.(string))
}

func (m maxHeap) Less(i,j int) bool{
	return m[i]<m[j]
}

func (m maxHeap) Len() int {
	return len(m)
}

func (m maxHeap) Swap(i,j int) {
	m[i],m[j]=m[j],m[i]
}

func main() {
	arr := maxHeap{"hello","how","are","you"}
	heap.Init(&arr)
	n := len(arr) -1
	for n>=0 {
		val := heap.Pop(&arr)
		fmt.Println(val)
		n--
	}
}