package main

import (
	"fmt"
	time2 "time"
)

func main() {
	values := []int{3, 2, 7, 8, 9}
	w := 15
	n := len(values)
	fmt.Println("Subset Sum Recursive")
	time := time2.Now().Nanosecond()
	fmt.Println(subsetsum(values, w, n))
	fmt.Printf("Time :- ")
	fmt.Println(time2.Now().Nanosecond() - time)

	fmt.Println("Subset Sum Memoization")
	dp = make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = make([]int, w+1)
		for j := 0; j < w+1; j++ {
			dp[i][j] = -1
		}
	}
	time = time2.Now().Nanosecond()
	fmt.Println(subsetsum_memoization(values, w, n))
	fmt.Printf("Time :- ")
	fmt.Println(time2.Now().Nanosecond() - time)

	fmt.Println("Subset Sum Top Down")
	time = time2.Now().Nanosecond()
	fmt.Println(subsetsum_topdown(values, w, n))
	fmt.Printf("Time :- ")
	fmt.Println(time2.Now().Nanosecond() - time)

	fmt.Println("\nEqual Sum Paritition\n")
	weight := []int{3, 7, 8, 9, 1}
	n = len(weight)
	sum := 0
	for i := 0; i < n; i++ {
		sum += weight[i]
	}
	fmt.Println(sum)
	if sum%2 != 0 {
		return
	}
	w = sum / 2

	fmt.Println("Equal Sum Paritition Recursive")
	time = time2.Now().Nanosecond()
	res := equalsumpartition_recursive(weight, w, n)
	fmt.Println(res)
	fmt.Printf("Time :- ")
	fmt.Println(time2.Now().Nanosecond() - time)

	fmt.Println("Equal Sum Paritition Memoized")
	dp = make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = make([]int, w+1)
		for j := 0; j < w+1; j++ {
			dp[i][j] = -1
		}
	}
	time = time2.Now().Nanosecond()
	res = equalsumpartition_memoize(weight, w, n)
	fmt.Println(res)
	fmt.Printf("Time :- ")
	fmt.Println(time2.Now().Nanosecond() - time)

	fmt.Println("Equal Sum Paritition Top Down")
	time = time2.Now().Nanosecond()
	res = equalsumpartition_topdown(weight, w, n)
	fmt.Println(res)
	fmt.Printf("Time :- ")
	fmt.Println(time2.Now().Nanosecond() - time)

}
