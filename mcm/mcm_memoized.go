package main

import (
	"math"
)

func mcm_memoized(nums []int, val ...int) int {
	i := 0
	j := 0
	if len(val) == 0 {
		i = 1
		j = len(nums) - 1
		if i >= j {
			return 0
		}
	} else {
		i = val[0]
		j = val[1]
		if i >= j {
			return 0
		}
	}
	if dp[i][j] != -1 {
		return dp[i][j]
	}
	min := math.MaxInt64
	for k := i; k <= j-1; k++ {
		tmpAns := mcm_memoized(nums, i, k) + mcm_memoized(nums, k+1, j) + nums[i-1]*nums[k]*nums[j]
		if tmpAns < min {
			min = tmpAns
		}
	}
	dp[i][j] = min
	return dp[i][j]
}
