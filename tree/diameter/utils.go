package main

import (
	"fmt"
	"strconv"
)

func (n *Node) addLeft(val int) {
	n.left = &Node{
		left:  nil,
		right: nil,
		val:   val,
	}
}

func (n *Node) addRight(val int) {
	n.right = &Node{
		left:  nil,
		right: nil,
		val:   val,
	}
}

func bfs(n *Node) {
	var queue []*Node
	visited := make(map[int]bool)
	visited[n.val]= true
	queue = append(queue,n)
	for len(queue)!=0 {
		length := len(queue) -1
		nodeOut := queue[length]
		fmt.Print(strconv.Itoa(nodeOut.val) + "->")
		queue = queue[:length]
		if nodeOut.left !=nil && !visited[nodeOut.left.val] {
			queue = append([]*Node{nodeOut.left},queue...)
		}
		if nodeOut.right !=nil && !visited[nodeOut.right.val] {
			queue = append([]*Node{nodeOut.right},queue...)
		}
	}
}