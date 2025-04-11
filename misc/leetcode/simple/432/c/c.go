package main

import (
	"container/heap"
	"math"
	"slices"
	"sort"
)

func minMaxWeight(n int, edges [][]int, _ int) int {
	if len(edges) < n-1 {
		return -1
	}

	type edge struct{ to, w int }
	g := make([][]edge, n)
	maxW := 0
	for _, e := range edges {
		x, y, w := e[0], e[1], e[2]
		g[y] = append(g[y], edge{x, w})
		maxW = max(maxW, w)
	}

	vis := make([]int, n)
	ans := 1 + sort.Search(maxW, func(upper int) bool {
		upper++
		left := n
		var dfs func(int)
		dfs = func(x int) {
			vis[x] = upper // 避免每次二分重新创建 vis 数组
			left--
			for _, e := range g[x] {
				if e.w <= upper && vis[e.to] != upper {
					dfs(e.to)
				}
			}
		}
		dfs(0)
		return left == 0
	})
	if ans > maxW {
		return -1
	}
	return ans
}

func minMaxWeight2(n int, edges [][]int, _ int) int {
	if len(edges) < n-1 {
		return -1
	}

	type edge struct{ to, w int }
	g := make([][]edge, n)
	for _, e := range edges {
		x, y, w := e[0], e[1], e[2]
		g[y] = append(g[y], edge{x, w})
	}

	dis := make([]int, n)
	for i := range dis {
		dis[i] = math.MaxInt
	}
	dis[0] = 0
	h := hp{{}}
	for len(h) > 0 {
		p := heap.Pop(&h).(pair)
		x := p.x
		d := p.dis
		if d > dis[x] {
			continue
		}
		for _, e := range g[x] {
			y := e.to
			newD := max(d, e.w)
			if newD < dis[y] {
				dis[y] = newD
				heap.Push(&h, pair{newD, y})
			}
		}
	}

	ans := slices.Max(dis)
	if ans == math.MaxInt {
		return -1
	}
	return ans
}

type pair struct{ dis, x int } // 路径最大边权, 节点编号
type hp []pair

func (h hp) Len() int           { return len(h) }
func (h hp) Less(i, j int) bool { return h[i].dis < h[j].dis }
func (h hp) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v any)        { *h = append(*h, v.(pair)) }
func (h *hp) Pop() (v any)      { a := *h; *h, v = a[:len(a)-1], a[len(a)-1]; return }
