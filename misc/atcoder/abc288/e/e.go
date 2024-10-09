package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// todo unPassed
func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, m, v int
	Fscan(in, &n, &m)
	a, c := make([]int, n), make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	for i := range c {
		Fscan(in, &c[i])
	}
	must := make(map[int]bool, m)
	for i := 0; i < m; i++ {
		Fscan(in, &v)
		must[v] = true
	}
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
		for j := range dp[i] {
			dp[i][j] = int(1e10)
		}
	}
	dp[0][0] = 0
	for i := 0; i < n; i++ {
		cost := int(1e9)
		for j := 0; j <= i; j++ {
			cost = min(cost, c[i-j])
			dp[i+1][j+1] = min(dp[i+1][j+1], dp[i][j]+a[i]+cost)
			if !must[i+1] {
				dp[i+1][j] = min(dp[i+1][j], dp[i][j])
			}
		}
	}
	var res = int(1e10)
	for i := m; i <= n; i++ {
		res = min(res, dp[n][i])
	}
	Fprintln(out, res)

}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() { run(os.Stdin, os.Stdout) }
