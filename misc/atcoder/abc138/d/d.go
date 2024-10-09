package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, q, x, y int
	Fscan(in, &n, &q)
	g, ans := make([][]int, n+1), make([]int, n+1)
	for i := 1; i < n; i++ {
		Fscan(in, &x, &y)
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}
	for i := 0; i < q; i++ {
		Fscan(in, &x, &y)
		ans[x] += y
	}
	var dfs func(x, fa int)
	dfs = func(x, fa int) {
		for _, to := range g[x] {
			if to != fa {
				ans[to] += ans[x]
				dfs(to, x)
			}
		}
	}
	dfs(1, 0)
	for _, v := range ans[1:] {
		Fprint(out, v, " ")
	}

}

func main() { run(os.Stdin, os.Stdout) }
