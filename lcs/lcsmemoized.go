package main

func lcsmemoized(x,y string,n, m int) int {
	if n==0 || m==0 {
		return 0
	}
	if dp[n][m]!=-1 {
		return dp[n][m]
	}
	if x[n-1]==y[m-1]{
		dp[n][m] = 1+lcsmemoized(x,y,n-1,m-1)
	} else {
		dp[n][m] = max(lcsmemoized(x,y,n-1,m),lcsmemoized(x,y,n,m-1))
	}
	return dp[n][m]
}