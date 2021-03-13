package main

import "fmt"

type Item struct {
	Val      int
	Priority int
}

type PQ []Item

var front, rear, capacity int

func (pq *PQ) Enqueue(val, priority int) {
	p := *pq
	p[rear] = Item{
		Val:      val,
		Priority: priority,
	}
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
	data := p[front]
	front = (front + 1) % capacity
	return data
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
}
