package main

import "fmt"

func main() {
	arr := []int {1,3,5,9,1,2,0,-1,4}
	n := len(arr)
	InsertionSort(arr,n)
	fmt.Println(arr)
}

func InsertionSort(arr []int,n int) {
	for i:=1;i<n;i++ {
		temp := arr[i]
		j:=i-1
		for j>=0 && arr[j]>temp{
			arr[j+1]=arr[j]
			j--
		}
		arr[j+1]=temp
	}
}
