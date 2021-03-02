package main

import "fmt"

func main() {
	arr := []int{2,4,6,-1,2,4,5,0,1,10}
	QuickSort(arr,0,len(arr)-1)
	fmt.Println(arr)
}

func QuickSort(arr []int,l,r int) {
	if l<r {
		pivot := Partition(arr,l,r)
		QuickSort(arr,l,pivot)
		QuickSort(arr,pivot+1,r)
	}
}

func Partition(arr []int,l,r int) int {
	pivot := arr[l]
	i:=l
	j:=r
	for i<j {
		for true {
			i++
			if i==r || !(arr[i]<=pivot) {
				break
			}
		}
		for true {
			j--
			if j==l || !(arr[j]>=pivot) {
				break
			}
		}
		if i<j {
			tmp:=arr[i]
			arr[i]=arr[j]
			arr[j]=tmp
		}
	}
	tmp := arr[l]
	arr[l] = arr[j]
	arr[j] = tmp
	return j
}