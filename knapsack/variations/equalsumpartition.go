package main

func equalsumpartition_recursive(weight []int, w, n int) bool {
	if w == 0 {
		return true
	}
	if w < 0 || n == 0 {
		return false
	}
	if weight[n-1] <= w {
		return equalsumpartition_recursive(weight, w-weight[n-1], n-1) || equalsumpartition_recursive(weight, w, n-1)
	} else {
		return equalsumpartition_recursive(weight, w, n-1)
	}
}