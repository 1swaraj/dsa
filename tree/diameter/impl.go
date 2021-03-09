package main

import (
	"fmt"
	"math"
)

func main() {
	root := &Node{
		left:  nil,
		right: nil,
		val:   1,
	}
	root.addLeft(2)
	root.addRight(3)
	root.left.addLeft(4)
	root.left.addRight(5)
	res = math.MinInt16
	bfs(root)
	fmt.Println("\nDiameter :- ")
	_ = diameter(root)
	fmt.Println(res)

	root = &Node{
		left:  nil,
		right: nil,
		val:   1,
	}
	root.addLeft(2)
	root.addRight(3)
	root.left.addLeft(4)
	root.left.addRight(5)
	root.right.addRight(6)
	res = math.MinInt16
	bfs(root)
	fmt.Println("\nDiameter :- ")
	_ = diameter(root)
	fmt.Println(res)

}
