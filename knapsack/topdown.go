package main

func topdown(val []int, weight []int, w, n int) int {
	dp = make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = make([]int, w+1)
	}
	for i:=0;i<n+1;i++ {
		for j:=0;j<w+1;j++ {
			if i == 0 || j == 0 {
				dp[i][j]=0
			} else if weight[i-1] <= j && val[i-1]+dp[i-1][j-weight[i-1]] >= dp[i-1][j] {
				dp[i][j] = val[i-1]+dp[i-1][j-weight[i-1]]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	return dp[n][w]
}
