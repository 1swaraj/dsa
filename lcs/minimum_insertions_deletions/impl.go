package main

import "fmt"

func main() {
	initialize(13,5,0)
	fmt.Println(minimum_insertions_deletions("geeksforgeeks", "geeks", 13, 5))
}