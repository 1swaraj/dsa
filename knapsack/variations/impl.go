package main

import (
	"fmt"
	time2 "time"
)

func main() {
	values := []int{3,2,7,8,9}
	w := 15
	n := len(values)
	fmt.Println("Subset Sum Recursive")
	time := time2.Now().Nanosecond()
	fmt.Println(subsetsum(values,w,n))
	fmt.Printf("Time :- ")
	fmt.Println(time2.Now().Nanosecond() - time)

	fmt.Println("Subset Sum Memoization")
	dp = make([][]int,n+1)
	for i:=0;i<n+1;i++ {
		dp[i] = make([]int,w+1)
		for j:=0;j<w+1;j++ {
			dp[i][j]=-1
		}
	}
	time = time2.Now().Nanosecond()
	fmt.Println(subsetsum_memoization(values,w,n))
	fmt.Printf("Time :- ")
	fmt.Println(time2.Now().Nanosecond() - time)

	fmt.Println("Subset Sum Top Down")
	time = time2.Now().Nanosecond()
	fmt.Println(subsetsum_topdown(values,w,n))
	fmt.Printf("Time :- ")
	fmt.Println(time2.Now().Nanosecond() - time)

}