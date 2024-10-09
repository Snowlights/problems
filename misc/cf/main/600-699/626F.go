//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"sort"
)

func CF626F(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	const mod = 1_000_000_007
	var n, k int
	Fscan(in, &n, &k)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	sort.Ints(a)
	dp := [200][101][1001]int{}
	for i := range dp {
		for j := range dp[i] {
			for k := range dp[i][j] {
				dp[i][j][k] = -1
			}
		}
	}
	var f func(int, int, int) int
	f = func(i, balance, leftK int) (res int) {
		if leftK < 0 || balance > i+1 {
			return
		}
		if i == 0 {
			if balance <= 1 {
				return 1
			}
			return
		}
		dv := &dp[i][balance][leftK]
		if *dv != -1 {
			return *dv
		}
		defer func() {
			*dv = res
		}()
		d := a[i] - a[i-1]
		res = f(i-1, balance+1, leftK-d*(balance+1))
		res = (res + (f(i-1, balance, leftK-d*balance))*(balance+1)) % mod
		if balance > 0 {
			res = (res + (f(i-1, balance-1, leftK-d*(balance-1)))*(balance)) % mod
		}
		return res
	}
	Fprint(out, f(n-1, 0, k))
}

func main() { CF626F(os.Stdin, os.Stdout) }
