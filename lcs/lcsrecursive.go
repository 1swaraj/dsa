package main

func lcsrecursive(x,y string,n, m int) int {
	if n==0 || m==0 {
		return 0
	}
	if x[n-1]==y[m-1]{
		return 1+lcsrecursive(x,y,n-1,m-1)
	} else {
		return max(lcsrecursive(x,y,n-1,m),lcsrecursive(x,y,n,m-1))
	}
}

func max(x,y int) int {
	if x>y {
		return x
	}
	return y
}