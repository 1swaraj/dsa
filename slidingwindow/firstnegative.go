package main

import "fmt"

func main() {
	arr := []int{100, 200, 300, 400}
	n := len(arr)
	k := 4
	sum := 0
	for i:=0;i<k;i++ {
		sum+=arr[i]
	}
	sumMax := sum
	for i:=k;i<n;i++ {
		sum = sum - arr[i-k] + arr[i]
		if sum > sumMax {
			sumMax = sum
		}
	}
	fmt.Println(sumMax)
}
