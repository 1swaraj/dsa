package main

import "fmt"

func main() {
	initialize(4,3,0)
	fmt.Println(shortest_common_subsequence("geek", "eke", 4, 3))
	initialize(6,7,0)
	fmt.Println(shortest_common_subsequence("AGGTAB", "GXTXAYB", 6, 7))
}