//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF2014F(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, c int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &c)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
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
		var dfs func(int, int) (int, int)
		dfs = func(v, fa int) (notDec, dec int) {
			dec = a[v]
			for _, w := range g[v] {
				if w != fa {
					nd, d := dfs(w, v)
					notDec += max(nd, d)
					dec += max(nd, d-c*2)
				}
			}
			return
		}
		Fprintln(out, max(dfs(0, -1)))
	}

}

func main() { CF2014F(os.Stdin, os.Stdout) }
