package main

import "fmt"

func main() {
	fmt.Println(lcsrecursive("ABCDGH", "AEDFHR", 6, 6))
	fmt.Println(lcsrecursive("AGGTAB", "GXTXAYB", 6, 7))
	initialize(6,6,-1)
	fmt.Println(lcsmemoized("ABCDGH", "AEDFHR", 6, 6))
	initialize(6,7,-1)
	fmt.Println(lcsmemoized("AGGTAB", "GXTXAYB", 6, 7))

	initialize(6,6,0)
	fmt.Println(lcstopdown("ABCDGH", "AEDFHR", 6, 6))
	initialize(6,7,0)
	fmt.Println(lcstopdown("AGGTAB", "GXTXAYB", 6, 7))

}