//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF369C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	type pair struct {
		to, t int
	}
	var n, from, to, t int
	Fscan(in, &n)
	g := make(map[int][]pair)
	ans := make([]int, 0, n)
	for i := 1; i < n; i++ {
		Fscan(in, &from, &to, &t)
		g[from] = append(g[from], pair{to, t})
		g[to] = append(g[to], pair{from, t})
	}

	var dfs func(fa, x int) bool
	dfs = func(fa, x int) (choose bool) {
		for _, to := range g[x] {
			if to.to == fa {
				continue
			}
			f := dfs(x, to.to)
			if !f && to.t == 2 {
				ans = append(ans, to.to)
				f = true
			}
			choose = choose || f
		}
		return
	}
	dfs(0, 1)
	Fprintln(out, len(ans))
	for _, v := range ans {
		Fprint(out, v, " ")
	}
}

func main() { CF369C(os.Stdin, os.Stdout) }
