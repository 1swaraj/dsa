package main

func countsubsetswithsum(weight []int, w, n int) int {
	if w == 0 {
		return 1
	}
	if w < 0 || n == 0 {
		return 0
	}
	if weight[n-1] <= w {
		return countsubsetswithsum(weight, w-weight[n-1], n-1) + countsubsetswithsum(weight, w, n-1)
	} else {
		return countsubsetswithsum(weight, w, n-1)
	}
}
