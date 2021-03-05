package main

import "math"

func memoization(val []int,weight []int,w,n int) int{
	if w==0 || n==0 {
		return 0
	}
	if dp[n][w] != -1 {
		return dp[n][w]
	}
	if weight[n]<=w{
		dp[n][w] = int(math.Max(float64(val[n]+memoization(val,weight,w-weight[n],n-1)),float64(memoization(val,weight,w,n-1))))
		return dp[n][w]
	} else {
		dp[n][w] = memoization(val,weight,w,n-1)
		return dp[n][w]
	}
}