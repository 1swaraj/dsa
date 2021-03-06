package main

func subsetsum_memoization(arr []int, w, n int) bool {
	if w == 0 {
		return true
	}
	if w < 0 || n == 0 {
		return false
	}
	if dp[n][w] == 0 {
		return false
	} else if dp[n][w] == 1 {
		return true
	}
	val := false
	if arr[n-1] <= w {
		val1 := subsetsum_memoization(arr, w-arr[n-1], n-1)
		val2 := subsetsum_memoization(arr, w, n-1)
		if val1 == true {
			val = true
		} else {
			val = val2
		}
	} else {
		val = subsetsum_memoization(arr, w, n-1)
	}
	if val==false {
		dp[n][w]=0
	} else {
		dp[n][w]=1
	}
	return val
}