//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF1547G(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, m int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m)
		g := make([][]int, n)
		for ; m > 0; m-- {
			var v, w int
			Fscan(in, &v, &w)
			g[v-1] = append(g[v-1], w-1)
		}

		ans, exit := make([]int, n), make([]bool, n)
		var dfs func(int, bool)
		dfs = func(v int, cycle bool) {
			exit[v] = true
			if cycle {
				ans[v] = -1
			} else if ans[v] < 2 {
				ans[v]++
			}
			for _, w := range g[v] {
				if ans[w] < 0 {
					continue
				}
				if cycle || exit[w] { // w 在环上
					dfs(w, true) // 从 w 出发能到达的点都在环上
				} else if ans[w] < 2 {
					dfs(w, false)
				}
			}
			exit[v] = false
		}
		dfs(0, false)
		for _, v := range ans {
			Fprint(out, v, " ")
		}
		Fprintln(out)
	}
}

func main() { CF1547G(os.Stdin, os.Stdout) }
