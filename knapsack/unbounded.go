package main

import (
	"math"
)

func unbounded(weight []int, w, n int) int {
	if w <= 0 || n<0 {
		return 0
	}
	if weight[n] <= w {
		return int(math.Max(float64(weight[n]+unbounded(weight, w-weight[n], n)), float64(unbounded(weight, w, n-1))))
	} else {
		return unbounded(weight, w, n-1)
	}
}

func unbounded_memoized(weight []int, w, n int) int {
	if w <= 0 || n<0 {
		return 0
	}
	if dp[n][w] != -1 {
		return dp[n][w]
	}
	if weight[n] <= w {
		dp[n][w]= int(math.Max(float64(weight[n]+unbounded(weight, w-weight[n], n)), float64(unbounded(weight, w, n-1))))
	} else {
		dp[n][w] = unbounded(weight, w, n-1)
	}
	return dp[n][w]
}

func unbounded_topdown(weight []int, w, n int) {
	dp = make([][]int,n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = make([]int, w+1)
	}
	for i:=0;i<n+1;i++{
		for j:=0;j<w+1;j++ {
			if i == 0 || j == 0 {
				dp[i][j]=0
			} else if weight[i-1] <= j && weight[i-1]+dp[i][j-weight[i-1]] >= dp[i-1][j] {
				dp[i][j] = weight[i-1]+dp[i][j-weight[i-1]]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}
}
