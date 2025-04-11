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

	var n, ans, diameter int
	Fscan(in, &n)
	type nb struct{ to, wt int }
	g := make([][]nb, n+1)
	for i := 1; i < n; i++ {
		var v, w, wt int
		Fscan(in, &v, &w, &wt)
		g[v] = append(g[v], nb{w, wt})
		g[w] = append(g[w], nb{v, wt})
		ans += wt
	}

	var dfs func(int, int) int
	dfs = func(v, fa int) (maxL int) {
		for _, e := range g[v] {
			if e.to != fa {
				subL := dfs(e.to, v) + e.wt
				diameter = max(diameter, maxL+subL)
				maxL = max(maxL, subL)
			}
		}
		return
	}
	dfs(1, 0)
	Fprint(out, ans*2-diameter)

}

func main() { run(os.Stdin, os.Stdout) }
