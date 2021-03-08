package main

import "fmt"

func main() {
	str := "geeksforgeeks"
	n := len(str)
	initialize(n,n,0)
	rev := ""
	for i:=n-1;i>=0;i-- {
		rev += string(str[i])
	}
	fmt.Println(minimum_to_palindrome(str,rev, n, n))
	str = "aebcbda"
	n = len(str)
	initialize(n,n,0)
	rev = ""
	for i:=n-1;i>=0;i-- {
		rev += string(str[i])
	}
	fmt.Println(minimum_to_palindrome(str,rev, n, n))
}