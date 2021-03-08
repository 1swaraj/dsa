package main

import "math"

func lcs_substring(x,y string,n, m int) int {
	if n==0 || m==0 {
		return 0
	}
	res := math.MinInt64
	for i:=1;i<n+1;i++ {
		for j:=1;j<m+1;j++ {
			if x[i-1] == y[j-1] {
				dp[i][j]=1+dp[i-1][j-1]
			} else {
				dp[i][j]=0
			}
			if dp[i][j]>res {
				res = dp[i][j]
			}
		}
	}
	return res
}

