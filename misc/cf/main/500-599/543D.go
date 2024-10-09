//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF543D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	const mod int = 1e9 + 7
	pow := func(x, n int) int {
		res := 1
		for ; n > 0; n >>= 1 {
			if n&1 > 0 {
				res = res * x % mod
			}
			x = x * x % mod
		}
		return res
	}

	var n, v int
	Fscan(in, &n)
	g := make([][]int, n)
	for w := 1; w < n; w++ {
		Fscan(in, &v)
		v--
		g[v] = append(g[v], w)
	}

	dp := make([]int, n)
	ex := make([]int, n)
	var f func(int)
	f = func(v int) {
		z := false
		dp[v] = 1
		ex[v] = 1
		for _, w := range g[v] {
			f(w)
			dw := dp[w] + 1
			dp[v] = dp[v] * dw % mod
			if z || dw != mod {
				ex[v] = ex[v] * dw % mod
			} else {
				z = true
			}
		}
	}
	f(0)

	ans := make([]int, n)
	var reRoot func(int, int)
	reRoot = func(v, dpFa int) {
		ans[v] = dp[v] * (dpFa + 1) % mod
		for _, w := range g[v] {
			df := 0
			if dp[w]+1 == mod {
				df = ex[v] * (dpFa + 1) % mod
			} else {
				df = ans[v] * pow(dp[w]+1, mod-2) % mod
			}
			reRoot(w, df)
		}
	}
	reRoot(0, 0)
	for _, v := range ans {
		Fprint(out, v, " ")
	}
}

func main() { CF543D(os.Stdin, os.Stdout) }
