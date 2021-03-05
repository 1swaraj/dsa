package main

import "math"

func recursive(value []int, weight []int, w, n int) int {
	if n == 0 || w == 0 {
		return 0
	}
	if weight[n] <= w {
		return int(math.Max(float64(value[n]+recursive(value, weight, w-weight[n], n-1)), float64(recursive(value, weight, w, n-1))))
	} else {
		return recursive(value, weight, w, n-1)
	}
}
