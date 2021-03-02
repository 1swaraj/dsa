package main

import "fmt"

func main() {
	arr := []int{2, 4, 6, -1, 2, 4, 5, 0, 1, 10, 12, 3, -2, -4}
	MergeSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}

func MergeSort(arr []int, l, r int) {
	if l < r {
		mid := (l + r) / 2
		MergeSort(arr, l, mid)
		MergeSort(arr, mid+1, r)
		Merge(arr, l, mid, r)
	}
}

func Merge(arr []int, l, m, r int) {
	var sorted []int
	i := l
	j := m + 1
	for i <= m && j <= r {
		if arr[i] < arr[j] {
			sorted = append(sorted, arr[i])
			i++
		} else {
			sorted = append(sorted, arr[j])
			j++
		}
	}
	for i <= m {
		sorted = append(sorted, arr[i])
		i++
	}
	for j <= r {
		sorted = append(sorted, arr[j])
		j++
	}
	k := len(sorted) - 1
	for x := r; x >= l; x-- {
		arr[x] = sorted[k]
		k--
	}
}
