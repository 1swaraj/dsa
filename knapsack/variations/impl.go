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

	fmt.Println("\nCount of Subsets with a Given Sum\n")
	weight = []int{3, 7, 8, 9, 1, 2, 5, 4}
	n = len(weight)
	w = 11

	fmt.Println("Count of Subsets with a Given Sum Recursive")
	time = time2.Now().Nanosecond()
	fmt.Println(countsubsetswithsum(weight, w, n))
	fmt.Printf("Time :- ")
	fmt.Println(time2.Now().Nanosecond() - time)

	fmt.Println("Count of Subsets with a Given Sum Memoized")
	dp = make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = make([]int, w+1)
		for j := 0; j < w+1; j++ {
			dp[i][j] = -1
		}
	}
	time = time2.Now().Nanosecond()
	fmt.Println(countsubsetswithsum_memoized(weight, w, n))
	fmt.Printf("Time :- ")
	fmt.Println(time2.Now().Nanosecond() - time)

	fmt.Println("Count of Subsets with a Given Sum Top Down")
	time = time2.Now().Nanosecond()
	fmt.Println(countsubsetswithsum_topdown(weight, w, n))
	fmt.Printf("Time :- ")
	fmt.Println(time2.Now().Nanosecond() - time)

	fmt.Println("\nMininum Subsets Sum Difference")
	weight = []int{7, 6, 11, 1}
	n = len(weight)
	time = time2.Now().Nanosecond()
	fmt.Println(minsubsetsumdiff(weight, n))
	fmt.Printf("Time :- ")
	fmt.Println(time2.Now().Nanosecond() - time)

	fmt.Println("\nCount the number of subset with a given difference")
	weight = []int{1, 2, 3, 3, 2}
	n = len(weight)
	d := 1
	time = time2.Now().Nanosecond()
	fmt.Println(countsubsetswithdiff(weight, d, n))
	fmt.Printf("Time :- ")
	fmt.Println(time2.Now().Nanosecond() - time)

	fmt.Println("\nTarget Sum")
	weight = []int{1,2,7,9,2}
	/*
	9,7 = 16
	 */
	n = len(weight)
	S := 17
	time = time2.Now().Nanosecond()
	fmt.Println(findTargetSumWays(weight, S))
	fmt.Printf("Time :- ")
	fmt.Println(time2.Now().Nanosecond() - time)

	fmt.Println("\nMinimum Cost to Cut a Stick")
	time = time2.Now().Nanosecond()
	fmt.Println(minCost(7, []int{1, 3, 4, 5}))
	fmt.Printf("Time :- ")
	fmt.Println(time2.Now().Nanosecond() - time)
}
