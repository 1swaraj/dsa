package main

import "fmt"

func main() {
	initialize(6,6,0)
	fmt.Println(lcs_printing("ABCDGH", "AEDFHR", 6, 6))
	initialize(6,7,0)
	fmt.Println(lcs_printing("AGGTAB", "AGXGXAB", 6, 7))

}