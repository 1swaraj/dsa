package main

import (
	"fmt"
	"math"
)

type Queue []int

var front,rear int

func (q *Queue) Enqueue(x int) {
	if q.IsFull() {
		fmt.Println("Overflows")
		return
	}
	queue := *q
	rear = rear + 1
	queue[rear-1] = x
}

func (q *Queue) Dequeue() int {
	if q.IsEmpty() {
		fmt.Println("Underflows")
		return math.MinInt64
	}
	queue := *q
	front = front+1
	return queue[front-1]
}

func (q *Queue) IsEmpty() bool {
	if front < 0 || front>rear {
		return true
	}
	return false
}

func (q *Queue) IsFull() bool {
	if rear >= 10 {
		return true
	}
	return false
}

func main() {
	queue := make(Queue,10)
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	queue.Enqueue(4)
	queue.Enqueue(5)
	queue.Enqueue(6)
	queue.Enqueue(7)
	queue.Enqueue(8)
	queue.Enqueue(9)
	queue.Enqueue(10)
	fmt.Println(queue.Dequeue())
	queue.Enqueue(7)
	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Dequeue())
}


