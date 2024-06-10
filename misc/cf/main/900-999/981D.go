//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF981D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, k int
	Fscan(in, &n, &k)
	a, sum := make([]int, n), make([]int, n+1)
	for i := range a {
		Fscan(in, &a[i])
		sum[i+1] = sum[i] + a[i]
	}

	pos := func(m int) bool {
		dp := make([][]bool, k+1)
		dp[0] = make([]bool, n+1)
		dp[0][0] = true
		for i := 1; i <= k; i++ {
			dp[i] = make([]bool, n+1)
			for used := i; used <= n; used++ {
				for prev := i - 1; prev < used; prev++ {
					dp[i][used] = dp[i-1][prev] && (sum[used]-sum[prev])&m == m
					if dp[i][used] {
						break
					}
				}
			}
		}
		return dp[k][n]
	}

	cur := 0
	for bit := 1 << 60; bit > 0; bit /= 2 {
		if pos(cur | bit) {
			cur |= bit
		}
	}
	Fprintln(out, cur)
}

func main() { CF981D(os.Stdin, os.Stdout) }
