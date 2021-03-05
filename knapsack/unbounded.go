package main

import (
	"math"
)

func unbounded(weight []int, w, n int) int {
	if n == 0 || w <= 0 {
		return w
	}
	if weight[n] <= w {
		return int(math.Max(float64(weight[n]+unbounded(weight, w-weight[n], n)), float64(unbounded(weight, w, n-1))))
	} else {
		return unbounded(weight, w, n-1)
	}
}
