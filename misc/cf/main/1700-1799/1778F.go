//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF1778F(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	const mx = 1000
	divisors := [mx + 1][]int{}
	for i := mx; i > 0; i-- { // 方便后面从大到小遍历因子
		for j := i; j <= mx; j += i {
			divisors[j] = append(divisors[j], i)
		}
	}

	// ceilSqrt[i]^2 是 i 的倍数
	ceilSqrt := [mx + 1]int{}
	for i := 1; i <= mx; i++ {
		ceilSqrt[i] = 1
		x := i
		for p := 2; p*p <= x; p++ {
			for p2 := p * p; x%p2 == 0; x /= p2 {
				ceilSqrt[i] *= p
			}
			if x%p == 0 {
				ceilSqrt[i] *= p
				x /= p
			}
		}
		if x > 1 {
			ceilSqrt[i] *= x
		}
	}

	var T, n, k int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &k)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		g := make([][]int, n)
		for i := 1; i < n; i++ {
			var v, w int
			Fscan(in, &v, &w)
			v--
			w--
			g[v] = append(g[v], w)
			g[w] = append(g[w], v)
		}
		if k == 0 {
			Fprintln(out, a[0])
			continue
		}

		subGcd := make([]int, n)
		var dfs0 func(int, int)
		dfs0 = func(v, fa int) {
			subGcd[v] = a[v]
			for _, w := range g[v] {
				if w != fa {
					dfs0(w, v)
					subGcd[v] = gcd(subGcd[v], subGcd[w])
				}
			}
		}
		dfs0(0, -1)

		var cnt int
		var dfs func(int, int, int)
		dfs = func(v, fa, d int) {
			if subGcd[v]%d == 0 {
				// 无需操作
				return
			}
			if subGcd[v]*subGcd[v]%d == 0 {
				cnt++ // 操作 v
				return
			}
			if len(g[v]) == 1 || a[v]*a[v]%d > 0 {
				cnt = 1e9 // 无法操作
				return
			}
			for _, w := range g[v] {
				if w != fa {
					dfs(w, v, ceilSqrt[d])
				}
			}
			cnt++ // 操作 v
		}

		for _, d := range divisors[a[0]] {
			cnt = 0
			for _, v := range g[0] {
				dfs(v, 0, d)
			}
			if cnt < k {
				Fprintln(out, a[0]*d)
				break
			}
		}
	}

}

func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}

func main() { CF1778F(os.Stdin, os.Stdout) }
