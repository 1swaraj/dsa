package main

import "fmt"

type Item struct {
	Val      int
	Priority int
}

type PQ []Item

var front, rear, capacity,count int

func (pq *PQ) Enqueue(val, priority int) {
	p := *pq
	p[rear] = Item{
		Val:      val,
		Priority: priority,
	}
	count ++
	p.Heapify(rear)
	rear = (rear + 1) % capacity

}

func (pq *PQ) Heapify(i int) {
	p := *pq
	parent := (i - 1) / 2
	if i == 1 || i == 2 {
		parent = 0
	} else if parent == 0 {
		return
	}
	if p[i].Priority > p[parent].Priority {
		p[parent], p[i] = p[i], p[parent]
		p.Heapify(parent)
	}
}

func (pq *PQ) Dequeue() Item {
	p := *pq
	count --
	p[0],p[count] = p[count],p[0]
	p.HeapifyDel(0)
	front = (front + 1) % capacity
	return p[count]
}

func (pq *PQ) HeapifyDel(i int) {
	p := *pq
	largest := i
	lchild := 2*i +1
	rchild := 2*i +2
	if lchild < count && p[lchild].Priority>p[largest].Priority {
		largest = lchild
	}
	if rchild < count && p[rchild].Priority > p[lchild].Priority {
		largest = rchild
	}
	if largest != i {
		p[largest],p[i] = p[i],p[largest]
		p.HeapifyDel(largest)
	}
}

func main() {
	capacity = 5
	pq := make(PQ, capacity)
	rear = 0
	front = 0
	pq.Enqueue(1, 2)
	fmt.Println(pq)
	pq.Enqueue(2, 3)
	fmt.Println(pq)
	pq.Enqueue(4, 4)
	fmt.Println(pq)
	pq.Enqueue(5, 7)
	fmt.Println(pq)
	pq.Enqueue(9, 5)
	fmt.Println(pq)
	fmt.Println(pq.Dequeue())
	fmt.Println(pq.Dequeue())
	fmt.Println(pq.Dequeue())
	fmt.Println(pq.Dequeue())
	fmt.Println(pq.Dequeue())
	pq.Enqueue(9, 5)
	pq.Enqueue(5, 7)
	fmt.Println(pq.Dequeue())
	fmt.Println(pq)

}
