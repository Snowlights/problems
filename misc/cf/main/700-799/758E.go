//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF758E(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n int
	Fscan(in, &n)
	type edge struct{ v, w, wt, p int }
	es := make([]edge, n-1)
	type nb struct{ to, i int }
	g := make([][]nb, n+1)
	for i := range es {
		var v, w, wt, p int
		Fscan(in, &v, &w, &wt, &p)
		es[i] = edge{v, w, wt, p}
		g[v] = append(g[v], nb{w, i})
		g[w] = append(g[w], nb{v, i})
	}

	a := make([]struct{ minSum, extraDecMin int }, n+1)
	var dfs func(int, int) (int, int)
	dfs = func(v, fa int) (minSum, maxSum int) {
		for _, e := range g[v] {
			w := e.to
			if w == fa {
				continue
			}
			mn, mx := dfs(w, v)
			p := es[e.i].p
			if mn < 0 || p < mn {
				return -1, 0
			}
			wt := es[e.i].wt
			minSum += max(wt-(p-mn), 1) + mn
			maxSum += wt + min(mx, p)
			a[w].extraDecMin = max(mx-p, 0)
		}
		a[v].minSum = minSum
		return
	}
	minSum, _ := dfs(1, 0)
	if minSum < 0 {
		Fprint(out, -1)
		return
	}

	dec := 0
	var modify func(int, int)
	modify = func(v, fa int) {
		for _, ew := range g[v] {
			w := ew.to
			if w == fa {
				continue
			}
			dec += a[w].extraDecMin
			modify(w, v)
			e := &es[ew.i]
			d := min(e.wt-1, e.p-a[w].minSum, dec)
			e.wt -= d
			e.p -= d
			dec -= d
		}
	}
	modify(1, 0)

	Fprintln(out, n)
	for _, e := range es {
		Fprintln(out, e.v, e.w, e.wt, e.p)
	}

}

func main() { CF758E(os.Stdin, os.Stdout) }
