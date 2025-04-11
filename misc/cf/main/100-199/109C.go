//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF109C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, v, w, sz, ans int
	var s string
	Fscan(in, &n)
	type nb struct {
		to    int
		lucky bool
	}
	g := make([][]nb, n)
	for i := 1; i < n; i++ {
		Fscan(in, &v, &w, &s)
		v--
		w--
		lucky := strings.Count(s, "4")+strings.Count(s, "7") == len(s)
		g[v] = append(g[v], nb{w, lucky})
		g[w] = append(g[w], nb{v, lucky})
	}

	vis := make([]bool, n)
	var dfs func(int)
	dfs = func(v int) {
		vis[v] = true
		sz++
		for _, e := range g[v] {
			if !e.lucky && !vis[e.to] {
				dfs(e.to)
			}
		}
	}
	for i, b := range vis {
		if !b {
			sz = 0
			dfs(i)
			ans += sz * (n - sz) * (n - sz - 1)
		}
	}
	Fprint(out, ans)

}

func main() { CF109C(os.Stdin, os.Stdout) }
