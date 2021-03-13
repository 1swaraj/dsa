package main

import "fmt"

func main() {
	arr := []int{1,3,0,-1,2,4,5}
	BubbleSort(arr,len(arr))
	fmt.Println(arr)
}

func BubbleSort(arr []int,n int) {
	for i:=0;i<n;i++ {
		for j:=n-i-1;j>=0;j-- {
			if arr[i]<arr[j]{
				arr[i],arr[j] = arr[j],arr[i]
			}
		}
	}
}