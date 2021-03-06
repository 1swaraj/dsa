package main

func countsubsetswithsum_memoized(weight []int, w, n int) int {
	if w == 0 {
		return 1
	}
	if w < 0 || n == 0 {
		return 0
	}
	if dp[n][w] != -1 {
		return dp[n][w]
	}
	if weight[n-1] <= w {
		dp[n][w] = countsubsetswithsum_memoized(weight, w-weight[n-1], n-1) + countsubsetswithsum_memoized(weight, w, n-1)
	} else {
		dp[n][w] = countsubsetswithsum_memoized(weight, w, n-1)
	}
	return dp[n][w]
}
