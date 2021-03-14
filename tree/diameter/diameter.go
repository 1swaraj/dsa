package main

import (
	"fmt"
	"math"
)

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

type Tree struct {
	Root *Node
}

func (n *Node) InsertLeft(val int) {
	n.Left = &Node{
		Val:   val,
		Left:  nil,
		Right: nil,
	}
}

func (n *Node) InsertRight(val int) {
	n.Right = &Node{
		Val:   val,
		Left:  nil,
		Right: nil,
	}
}

func Display(node *Node) {
	if node != nil {
		fmt.Printf("%d ",node.Val)
		Display(node.Left)
		Display(node.Right)
	}
}

var res int
var ans int

func Diameter(node *Node) int {
	if node == nil {
		return 0
	}
	res = 1 + Height(node)
	return res
}

func Height(node *Node) int {
	if node==nil {
		return 0
	}
	return 1 + max(Height(node.Left),Height(node.Right))
}

func max(x,y int) int {
	if x>y {
		return x
	}
	return y
}

func main() {
	var tree Tree
	tree.Root= &Node{
		Val:1,
	}
	tree.Root.InsertLeft(2)
	tree.Root.InsertRight(3)
	tree.Root.Left.InsertLeft(4)
	tree.Root.Left.InsertRight(5)
	Display(tree.Root)
	ans = math.MaxInt64
	fmt.Println()
	fmt.Println(Diameter(tree.Root))
}
