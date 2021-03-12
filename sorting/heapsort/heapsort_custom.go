package main

import "fmt"

func main() {
	arr := []int{2,4,6,-1,2,4,5,0,1,11,10}
	n:= len(arr)
	for i:=n/2;i>=0;i-- {
		heapify(arr,n,i)
	}
	fmt.Println("Heap :- ")
	fmt.Println(arr)
	fmt.Println("Heap Sort :- ")
	for i:=0;i<n;i++ {
		pop := arr[0]
		arr[0] = arr[n-i-1]
		arr[n-i-1] = pop
		heapify(arr,n-i-1,0)
	}
	fmt.Println(arr)
}

func heapify(arr []int,n,i int) {
	largest := i
	lchild := 2*i + 1
	rchild := 2*i + 2
	if lchild < n && arr[lchild] > arr[largest] {
		largest = lchild
	}
	if rchild < n && arr[rchild] > arr[largest] {
		largest = rchild
	}
	if largest != i {
		arr[i],arr[largest] = arr[largest],arr[i]
		heapify(arr,n,largest)
	}
}