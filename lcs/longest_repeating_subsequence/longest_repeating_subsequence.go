package main

func longest_repeating_subsequence(x, y string, n int) int {
	for i := 1; i < n+1; i++ {
		for j := 1; j < n+1; j++ {
			if i!=j && x[i-1] == y[j-1] {
				dp[i][j] = 1 + dp[i-1][j-1]
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[n][n]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
