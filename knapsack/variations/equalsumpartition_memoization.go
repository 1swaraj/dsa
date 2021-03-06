package main

func equalsumpartition_memoize(weight []int, w, n int) bool {
	if w == 0{
		return true
	} else if w < 0 || n==0 {
		return false
	}
	if dp[n][w] != -1 {
		if dp[n][w] == 1 {
			return true
		} else {
			return false
		}
	}
	if weight[n-1] <= w {
		return equalsumpartition_memoize(weight,w,n-1) || equalsumpartition_memoize(weight,w-weight[n-1],n-1)
	} else {
		return equalsumpartition_memoize(weight,w,n-1)
	}
}
