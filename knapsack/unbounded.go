package main

import (
	"math"
)

func unbounded(weight []int, w, n int) int {
	if w <= 0 || n<0 {
		return 0
	}
	if weight[n] <= w {
		return int(math.Max(float64(weight[n]+unbounded(weight, w-weight[n], n)), float64(unbounded(weight, w, n-1))))
	} else {
		return unbounded(weight, w, n-1)
	}
}
