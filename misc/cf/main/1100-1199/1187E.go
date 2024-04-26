//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF1187E(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, ans int
	Fscan(in, &n)
	g := make([][]int, n+1)
	for i := 1; i < n; i++ {
		var v, w int
		Fscan(in, &v, &w)
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}

	size := make([]int, n+1)
	var dfs func(int, int) int
	dfs = func(v, fa int) (sum int) {
		size[v] = 1
		for _, w := range g[v] {
			if w != fa {
				sum += dfs(w, v)
				size[v] += size[w]
			}
		}
		return sum + size[v]
	}

	var reroot func(int, int, int)
	reroot = func(v, fa, res int) {
		ans = max(ans, res)
		for _, w := range g[v] {
			if w != fa {
				reroot(w, v, res+n-size[w]*2)
			}
		}
	}
	reroot(1, 0, dfs(1, 0))
	Fprint(out, ans)
	
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() { CF1187E(os.Stdin, os.Stdout) }
