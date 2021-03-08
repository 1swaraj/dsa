package main

import "fmt"

func main() {
	initialize(8, 8, 0)
	fmt.Println(longest_repeating_subsequence("AABEBCDD", "AABEBCDD", 8))
}
