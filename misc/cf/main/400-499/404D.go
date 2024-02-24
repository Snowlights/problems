//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF404D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var s string
	mod := int(1e9 + 7)
	Fscan(in, &s)
	n := len(s)
	dp := make([][3]int, n)
	for i := range dp {
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}
	var f func(int, int) int
	f = func(i, j int) (res int) {
		if i < 0 {
			if j == 1 {
				return 0
			}
			return 1
		}
		dv := &dp[i][j]
		if *dv != -1 {
			return *dv
		}
		defer func() { *dv = res }()
		b := s[i]
		if j == 0 {
			switch b {
			case '0':
				return f(i-1, 0)
			case '1':
				return f(i-1, 1)
			case '2', '*':
				return 0
			default:
				return (f(i-1, 0) + f(i-1, 1)) % mod
			}
		} else if j == 1 {
			switch b {
			case '0', '1', '2':
				return 0
			default:
				return f(i-1, 2)
			}
		} else { // *
			switch b {
			case '0':
				return 0
			case '1':
				return f(i-1, 0)
			case '2':
				return f(i-1, 1)
			case '*':
				return f(i-1, 2)
			default:
				return (f(i-1, 0) + f(i-1, 1) + f(i-1, 2)) % mod
			}
		}
	}
	Fprint(out, (f(n-1, 0)+f(n-1, 1))%mod)

}

func main() { CF404D(os.Stdin, os.Stdout) }
