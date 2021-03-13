package main

type Node struct {
	Val int
	Next *Node
	Visited bool
}

type LinkedList struct {
	Head *Node
}


