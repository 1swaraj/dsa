package main

type Node struct {
	Left   *Node
	Right  *Node
	Key    int
	Height int
}

type Tree struct {
	Head *Node
}