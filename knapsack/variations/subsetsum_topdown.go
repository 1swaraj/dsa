package main

func subsetsum_topdown(arr []int, w, n int) bool {
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
			if arr[i-1] <= j {
				dp[i][j] = dp[i-1][j-arr[i-1]] || dp[i-1][j]
			} else {
				dp[i][j]=dp[i-1][j]
			}
		}
	}
	return dp[n][w]
}