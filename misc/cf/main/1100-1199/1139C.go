//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF1139C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	const mod int64 = 1e9 + 7
	var n, k, u, v, w int64
	Fscan(in, &n, &k)
	g := make([][]int64, n)
	for i := int64(1); i < n; i++ {
		Fscan(in, &u, &v, &w)
		if w == 0 {
			u--
			v--
			g[u] = append(g[u], v)
			g[v] = append(g[v], u)
		}
	}
	pow := func(x, n int64) int64 {
		//x %= mod
		res := int64(1)
		for ; n > 0; n >>= 1 {
			if n&1 == 1 {
				res = res * x % mod
			}
			x = x * x % mod
		}
		return res
	}

	vis := make([]bool, n)
	var find func(int) int
	find = func(from int) int {
		res := 1
		vis[from] = true
		for _, to := range g[from] {
			if !vis[to] {
				res += find(int(to))
			}
		}
		return res
	}

	ans := pow(n, k)
	for from, v := range vis {
		if !v {
			ans -= pow(int64(find(from)), k)
		}
	}
	Fprintln(out, (ans%mod+mod)%mod)
}

func main() { CF1139C(os.Stdin, os.Stdout) }
