package main

func equalsumpartition_topdown(weight []int, w, n int) bool {
	t := make([][]bool,n+1)
	for i:=0;i<n+1;i++ {
		t[i] = make([]bool,w+1)
		for j:=0;j<w+1;j++{
			if i==0 {
				t[i][j]=false
			} else if j==0 {
				t[i][j]=true
			}
		}
	}
	t[0][0]=true
	for i:=1;i<n+1;i++ {
		for j:=1;j<w+1;j++ {
			if weight[i-1]<=j {
				t[i][j] = t[i-1][j-weight[i-1]] || t[i-1][j]
			} else {
				t[i][j] = t[i-1][j]
			}
		}
	}
	return t[n][w]
}
