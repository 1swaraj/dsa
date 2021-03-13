package main

import "fmt"

func (l *LinkedList) InsertAtEnd(node *Node) {
	head := l.Head
	for head.Next != nil {
		head = head.Next
	}
	head.Next = node
}

func (l *LinkedList) InsertAtBeginning(node *Node) {
	head := l.Head
	node.Next = head
	l.Head = node
}

func (l *LinkedList) InsertAtPos(node *Node, n int) {
	head := l.Head
	for head.Next != nil && n > 0 {
		n--
		head = head.Next
	}
	node.Next = head.Next
	head.Next = node
}

func (l *LinkedList) DeleteAtEnd() {
	head := l.Head
	for head.Next.Next != nil {
		head = head.Next
	}
	head.Next = nil
}

func (l *LinkedList) DeleteAtBeginning() {
	head := l.Head.Next
	l.Head = head
}

func (l *LinkedList) DeleteAtPos(n int) {
	head := l.Head
	for head.Next.Next != nil && n > 0 {
		n--
		head = head.Next
	}
	head.Next = head.Next.Next
}

func (l *LinkedList) Print() {
	head := l.Head
	for head != nil {
		fmt.Printf("%d -> ", head.Val)
		head = head.Next
	}
	fmt.Println()
}

func (l *LinkedList) Search(val int) bool {
	head := l.Head
	for head != nil {
		if head.Val == val {
			return true
		}
		head = head.Next
	}
	return false
}

func (l *LinkedList) Reverse() {
	head := l.Head
	var prev *Node
	for head !=nil {
		next := head.Next
		head.Next = prev
		prev = head
		head = next
	}
	l.Head = prev
}
