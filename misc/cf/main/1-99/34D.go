//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF34D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, r1, r2, v int
	Fscan(in, &n, &r1, &r2)
	g := make([][]int, n+1)
	for w := 1; w <= n; w++ {
		if w != r1 {
			Fscan(in, &v)
			g[v] = append(g[v], w)
			g[w] = append(g[w], v)
		}
	}

	ans := make([]int, n+1)
	var dfs func(int, int)
	dfs = func(v, fa int) {
		ans[v] = fa
		for _, w := range g[v] {
			if w != fa {
				dfs(w, v)
			}
		}
	}
	dfs(r2, 0)
	for i := 1; i <= n; i++ {
		if i != r2 {
			Fprint(out, ans[i], " ")
		}
	}

}

func main() { CF34D(os.Stdin, os.Stdout) }
