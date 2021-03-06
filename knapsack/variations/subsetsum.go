package main

func subsetsum(arr []int, w, n int) bool {
	if w == 0 {
		return true
	}
	if w < 0 || n == 0 {
		return false
	}
	if arr[n-1] <= w {
		return subsetsum(arr, w-arr[n-1], n-1) || subsetsum(arr, w, n-1)
	} else {
		return subsetsum(arr, w, n-1)
	}
}