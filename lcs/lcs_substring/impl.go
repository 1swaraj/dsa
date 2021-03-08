package main

import "fmt"

func main() {
	initialize(6,6,0)
	fmt.Println(lcs_substring("ABCDGH", "AEDFHR", 6, 6))
	initialize(6,7,0)
	fmt.Println(lcs_substring("AGGTAB", "AGXGXAB", 6, 7))

}