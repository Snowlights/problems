//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF659E(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, m, v, w, ans, V, E int
	Fscan(in, &n, &m)
	g := make([][]int, n)
	for ; m > 0; m-- {
		Fscan(in, &v, &w)
		v--
		w--
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}
	vis := make([]bool, n)
	var dfs func(int)
	dfs = func(v int) {
		vis[v] = true
		V++
		E += len(g[v])
		for _, w := range g[v] {
			if !vis[w] {
				dfs(w)
			}
		}
	}
	for i, b := range vis {
		if !b {
			V, E = 0, 0
			dfs(i)
			if V == E/2+1 {
				ans++
			}
		}
	}
	Fprint(out, ans)
	
}

func main() { CF659E(os.Stdin, os.Stdout) }
