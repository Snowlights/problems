//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF570E(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	const mod int = 1e9 + 7

	var n, m int
	Fscan(in, &n, &m)
	s := make([]string, n+1)
	for i := 1; i <= n; i++ {
		Fscan(in, &s[i])
		s[i] = " " + s[i]
	}

	dp := make([][][2]int, n+2)
	for i := range dp {
		dp[i] = make([][2]int, m+2)
	}
	for xp := 1; xp <= n; xp++ {
		for x := n; x > 0; x-- {
			for y := m; y > 0; y-- {
				yp := n + m + 2 - x - y - xp
				if 1 > yp || yp > m || x > xp || y > yp {
					dp[x][y][xp%2] = 0
				} else if x+y == (n+m+2)/2 {
					if s[x][y] == s[xp][yp] {
						dp[x][y][xp%2] = 1
					}
				} else if s[x][y] == s[xp][yp] {
					dp[x][y][xp%2] = ((dp[x+1][y][xp%2]+dp[x+1][y][1-xp%2])%mod + (dp[x][y+1][xp%2]+dp[x][y+1][1-xp%2])%mod) % mod
				} else {
					dp[x][y][xp%2] = 0
				}
			}
		}
	}
	Println(dp)
	Fprintln(out, dp[1][1][n%2])
}

func main() { CF570E(os.Stdin, os.Stdout) }
