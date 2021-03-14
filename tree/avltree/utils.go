package main

import "fmt"

func Insert(node *Node, val int) *Node {
	if node == nil {
		return &Node{
			Left:   nil,
			Right:  nil,
			Key:    val,
			Height: 1,
		}
	}
	if val < node.Key {
		node.Left = Insert(node.Left, val)
	} else if val > node.Key {
		node.Right = Insert(node.Right, val)
	} else {
		return node
	}
	node.Height = 1 + max(height(node.Left), height(node.Right))
	balance := height(node.Left) - height(node.Right)
	if balance > 1 && val < node.Left.Key {
		return RotateRight(node)
	} else if balance < -1 && val > node.Right.Key {
		return RotateLeft(node)
	} else if balance > 1 && val > node.Left.Key {
		node.Left = RotateLeft(node.Left)
		return RotateRight(node)
	} else if balance < -1 && val < node.Right.Key {
		node.Right = RotateRight(node.Right)
		return RotateLeft(node)
	}

	return node
}

func RotateRight(node *Node) *Node {
	x := node.Left
	y := x.Right
	x.Right = node
	node.Left = y
	node.Height = max(height(node.Left), height(node.Right)) + 1
	x.Height = max(height(x.Left), height(x.Right)) + 1
	return x
}

func height(node *Node) int {
	if node == nil {
		return 0
	}
	return node.Height
}

func RotateLeft(node *Node) *Node {
	x := node.Right
	y := x.Left
	x.Left = node
	node.Right = y
	x.Height = 1 + max(height(x.Left), height(x.Right))
	node.Height = 1 + max(height(node.Left), height(node.Right))
	return x
}

func preOrder(node *Node) {
	if node != nil {
		fmt.Printf("%d ", node.Key)
		preOrder(node.Left)
		preOrder(node.Right)
	}
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}
