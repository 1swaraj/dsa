package main

func longest_palindrome_printing(x,y string,n, m int) string {
	if n==0 || m==0 {
		return ""
	}
	res := ""
	for i:=1;i<n+1;i++ {
		for j:=1;j<m+1;j++ {
			if x[i-1] == y[j-1] {
				dp[i][j]=1+dp[i-1][j-1]
			} else {
				dp[i][j]=dp[i-1][j-1]
			}
		}
	}
	for n>0 && m>0 {
		if x[n-1]==y[m-1] {
			res += string(x[n-1])
			n = n - 1
			m = m - 1
			continue
		}
		if dp[n-1][m]>=dp[n][m-1] {
			n = n - 1
		} else {
			m = m - 1
		}
	}
	ans := ""
	for i:=len(res)-1;i>=0;i-- {
		ans+=string(res[i])
	}
	return ans
}

