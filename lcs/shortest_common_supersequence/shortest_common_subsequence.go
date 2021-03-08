package main

import (
	"strings"
)

func shortest_common_subsequence(x, y string, n, m int) string {
	if n == 0 || m == 0 {
		return ""
	}
	res := ""
	for i := 1; i < n+1; i++ {
		for j := 1; j < m+1; j++ {
			if x[i-1] == y[j-1] {
				dp[i][j] = 1 + dp[i-1][j-1]
			} else {
				dp[i][j] = dp[i-1][j-1]
			}
		}
	}
	for n > 0 && m > 0 {
		if x[n-1] == y[m-1] {
			res += string(x[n-1])
			n = n - 1
			m = m - 1
			continue
		}
		if dp[n-1][m] >= dp[n][m-1] {
			n = n - 1
		} else {
			m = m - 1
		}
	}
	ans := ""
	for i := len(res) - 1; i >= 0; i-- {
		ans += string(res[i])
	}
	i := 0
	j := 0
	k := 0
	final := ""
	n = len(x)
	m = len(y)
	for i < n && j < m && k < len(ans) {
		if (i+j)%2 == 0 {
			final += string(x[i])
			if x[i] == ans[k] {
				y = strings.Replace(y, string(ans[k]), "", 1)
				k++
			}
			i++
		} else {
			final += string(y[j])
			if y[j] == ans[k] {
				x = strings.Replace(x, string(ans[k]), "", 1)
				k++
			}
			j++
		}
		n = len(x)
		m = len(y)
	}
	for i < n {
		final += string(x[i])
		i++
	}
	for j < m {
		final += string(y[j])
		j++
	}
	return final
}
