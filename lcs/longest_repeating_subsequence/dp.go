package main

var dp [][]int

func initialize(n, m, v int) {
	dp = make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = make([]int, m+1)
		for j := 0; j < m+1; j++ {
			dp[i][j] = v
		}
	}
}
