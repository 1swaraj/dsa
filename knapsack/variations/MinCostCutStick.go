package main

/*
Link :- https://leetcode.com/contest/weekly-contest-201/problems/minimum-cost-to-cut-a-stick/
 */

import (
	"math"
	"sort"
)

func minCost(n int, cuts []int) int {
	sort.Ints(cuts)
	dp = make([][]int, 110)
	for i := 0; i < 110; i++ {
		dp[i] = make([]int, 110)
		for j := 0; j < 110; j++ {
			dp[i][j] = -1
		}
	}
	return findMinCost(cuts, 0, len(cuts)-1, 0, n)
}

func findMinCost(weight []int, i, j, l, r int) int {
	if i < 0 || j >= len(weight) || i > j {
		return 0
	}

	if i == j {
		return r - l
	}

	if dp[i][j] != -1 {
		return dp[i][j]
	}

	ans := math.MaxInt32
	for x := i; x <= j; x++ {
		ans = min(ans, (r-l)+findMinCost(weight, i, x-1, l, weight[x])+findMinCost(weight, x+1, j, weight[x], r))
	}
	dp[i][j] = ans
	return dp[i][j]
}

func min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}
