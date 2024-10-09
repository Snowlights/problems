package main

import "sort"

func maxIncreasingCells(a [][]int) (ans int) {

	type pair struct {
		x, y int
	}
	g := make(map[int][]pair)
	for i, val := range a {
		for j, v := range val {
			g[v] = append(g[v], pair{i, j})
		}
	}
	p := make([]int, 0, len(g))
	for k := range g {
		p = append(p, k)
	}
	sort.Ints(p)

	row, col := make([]int, len(a)), make([]int, len(a[0]))
	for _, v := range p {
		pos := g[v]
		mx := make([]int, len(pos))
		for i, vv := range pos {
			mx[i] = max(row[vv.x], col[vv.y]) + 1
			ans = max(ans, mx[i])
		}
		for i, vv := range pos {
			row[vv.x] = max(row[vv.x], mx[i])
			col[vv.y] = max(col[vv.y], mx[i])
		}
	}

	return
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
