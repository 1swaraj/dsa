package main

import "fmt"

func recusive(values,weight []int,w,n int) int {
	if n==0 || w==0 {
		return 0
	}
	return max(values[n]+recusive(values,weight,w-weight[n],n-1),recusive(values,weight,w,n-1))
}

func max(x,y int) int {
	if x>y {
		return x
	}
	return y
}

func main () {
	values := []int{60,100,120}
	weight := []int{10,20,30}
	W := 50
	n:=len(weight)-1
	fmt.Println(recusive(values,weight,W,n))
}
