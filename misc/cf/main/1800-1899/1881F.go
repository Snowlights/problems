//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF1881F(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, k, v, w int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &k)
		sp := make([]bool, n)
		for ; k > 0; k-- {
			Fscan(in, &v)
			v--
			sp[v] = true
		}
		rt := v
		g := make([][]int, n)
		for i := 1; i < n; i++ {
			Fscan(in, &v, &w)
			v--
			w--
			g[v] = append(g[v], w)
			g[w] = append(g[w], v)
		}

		diameter := 0
		var dfs func(int, int) int
		dfs = func(v, fa int) (maxL int) {
			for _, w := range g[v] {
				if w != fa {
					subL := dfs(w, v) + 1
					diameter = max(diameter, maxL+subL)
					maxL = max(maxL, subL)
				}
			}
			if maxL == 0 && !sp[v] {
				return -1
			}
			return
		}
		dfs(rt, -1)
		Fprintln(out, (diameter+1)/2)
	}

}

func main() { CF1881F(os.Stdin, os.Stdout) }
