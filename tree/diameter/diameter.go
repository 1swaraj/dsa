package main

func diameter(root *Node) int{
	if root == nil {
		return 0
	}

	l := diameter(root.left)

	r := diameter(root.right)

	temp := 1 + max(l,r)
	ans := 1 + l + r

	res = max(res,ans)

	return temp
}

func max(x,y int) int {
	if x>y {
		return x
	}
	return y
}