//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF1081C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, m, k int
	Fscan(in, &n, &m, &k)
	mod := int(998244353)
	// 前i个有j个色块不同的方案数
	//dp := make([][]int, n)
	//for i := range dp {
	//	dp[i] = make([]int, k+1)
	//	dp[i][0] = m
	//}
	//for i := 1; i < n; i++ {
	//	for j := 1; j <= k; j++ {
	//		dp[i][j] += dp[i-1][j] + dp[i-1][j-1]*(m-1)%mod
	//	}
	//}
	// Fprintln(out, dp[n-1][k]%mod)

	dp := make([]int, k+1)
	dp[0] = m

	for i := 1; i < n; i++ {
		for j := k; j > 0; j-- {
			dp[j] = dp[j] + dp[j-1]*(m-1)%mod
		}
	}

	Fprintln(out, dp[k]%mod)
}

func main() { CF1081C(os.Stdin, os.Stdout) }
