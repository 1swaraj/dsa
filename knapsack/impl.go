package main

import (
	"fmt"
	time2 "time"
)

func main() {
	values := []int{60,100,120}
	weight := []int{10,20,30}
	w := 50
	n := len(weight) - 1
	time := time2.Now().Nanosecond()
	max := recursive(values, weight, w, n)
	fmt.Println(max)
	fmt.Println(time2.Now().Nanosecond()-time)
	time = time2.Now().Nanosecond()
	dp = make([][]int,n+1)
	for i:=0;i<n+1;i++{
		dp[i] = make([]int,w+1)
		for j:=0;j<w+1;j++ {
			dp[i][j]=-1
		}
	}
	fmt.Println(memoization(values, weight, w, n))
	fmt.Println(time2.Now().Nanosecond()-time)
}
