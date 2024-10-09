package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"sort"
)

func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n int
	Fscan(in, &n)
	type pair struct {
		x, y, wt int
	}
	pL := make([]pair, n-1)
	for i := range pL {
		Fscan(in, &pL[i].x, &pL[i].y, &pL[i].wt)
	}
	sort.Slice(pL, func(i, j int) bool {
		return pL[i].wt < pL[j].wt
	})

	fa, size := make([]int, n+1), make([]int, n+1)
	for i := range fa {
		fa[i] = i
		size[i] = 1
	}
	var find func(x int) int
	find = func(x int) int {
		if fa[x] != x {
			fa[x] = find(fa[x])
		}
		return fa[x]
	}

	ans := 0
	for _, p := range pL {
		from, to := find(p.x), find(p.y)
		ans += size[from] * size[to] * p.wt
		fa[from] = to
		size[to] += size[from]
	}

	Fprintln(out, ans)

}

func main() { run(os.Stdin, os.Stdout) }
