//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF213B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	const mod = 1_000_000_007
	const mx = 101
	C := [mx][mx]int{}
	for i := 0; i < mx; i++ {
		C[i][0] = 1
		for j := 1; j <= i; j++ {
			C[i][j] = (C[i-1][j-1] + C[i-1][j]) % mod
		}
	}

	var n, ans int
	Fscan(in, &n)
	a := [10]int{}
	for i := range a {
		Fscan(in, &a[i])
	}

	memo := make([][10]int, n+1)
	for i := range memo {
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}
	var dfs func(int, int) int
	dfs = func(i, j int) int {
		if j == 9 {
			if i < a[9] {
				return 0
			}
			return 1
		}
		p := &memo[i][j]
		if *p != -1 {
			return *p
		}
		i2 := i
		if j == 0 {
			i2--
		}
		res := 0
		for k := a[j]; k <= i; k++ {
			res = (res + dfs(i-k, j+1)*C[i2][k]) % mod
		}
		*p = res
		return res
	}
	for i := 1; i <= n; i++ {
		ans += dfs(i, 0)
	}
	Fprint(out, ans%mod)

}

func main() { CF213B(os.Stdin, os.Stdout) }
