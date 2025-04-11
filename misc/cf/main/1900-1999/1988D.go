//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF1988D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n)
		s := 0
		for i := range a {
			Fscan(in, &a[i])
			s += a[i]
		}
		g := make([][]int, n)
		for range n - 1 {
			var v, w int
			Fscan(in, &v, &w)
			v--
			w--
			g[v] = append(g[v], w)
			g[w] = append(g[w], v)
		}
		var dfs func(int, int) (int, int, int)
		dfs = func(v, fa int) (mn, mn2, mnI int) {
			f := make([]int, min(20, len(g[v])+1))
			for _, w := range g[v] {
				if w == fa {
					continue
				}
				m, m2, i := dfs(w, v)
				for j := range f {
					if j != i {
						f[j] += m
					} else {
						f[j] += m2
					}
				}
			}
			mn = int(1e18)
			for i, res := range f {
				res += i * a[v]
				if res < mn {
					mn2 = mn
					mn = res
					mnI = i
				} else if res < mn2 {
					mn2 = res
				}
			}
			return
		}
		mn, _, _ := dfs(0, -1)
		Fprintln(out, s+mn)
	}

}

func main() { CF1988D(os.Stdin, os.Stdout) }
