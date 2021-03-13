package main

import "fmt"

func main() {
	arr := []int{1,35,7,0,-3,4,6}
	n := len(arr)
	SelectionSort(arr,n)
	fmt.Println(arr)
}

func SelectionSort(arr []int, n int) {
	for i:=0;i<n;i++ {
		min := arr[i]
		pos := i
		for j:=i+1;j<n;j++ {
			if arr[j]<min {
				min=arr[j]
				pos = j
			}
		}
		arr[i],arr[pos] = arr[pos],arr[i]
	}
}