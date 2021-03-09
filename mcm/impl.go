package main

import "fmt"

var dp[][]int
func main() {
	values := []int{40,20,30,10,30}
	n := len(values)
	fmt.Println(mcm(values))
	dp = make([][]int,n+1)
	for i:=0;i<n+1;i++ {
		dp[i]=make([]int,n+1)
		for j:=0;j<n+1;j++ {
			dp[i][j]=-1
		}
	}
	fmt.Println(mcm_memoized(values))
}