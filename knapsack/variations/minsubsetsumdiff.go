package main

import (
	"math"
)

func minsubsetsumdiff(weight []int, n int) int {
	w := 0
	for _, val := range weight {
		w += val
	}
	dp := make([][]bool,n+1)
	for i:=0;i<n+1;i++ {
		dp[i] = make([]bool,w+1)
		dp[i][0]=true
	}
	for j:=0;j<w+1;j ++ {
		dp[0][j]=false
	}
	dp[0][0]=true
	for i:=1;i<n+1;i++{
		for j:=1;j<w+1;j++ {
			if weight[i-1] <= j {
				dp[i][j] = dp[i-1][j-weight[i-1]] || dp[i-1][j]
			} else {
				dp[i][j]=dp[i-1][j]
			}
		}
	}
	diff := math.MaxInt64
	half := len(dp[n])/2
	for i,val := range dp[n][:half] {
		if val {
			if w-2*i<diff{
				diff=w-2*i
			}
		}
	}
	return diff
}