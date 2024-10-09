package main

import "container/heap"

func modifiedGraphEdges(n int, edges [][]int, st, end, target int) (ans [][]int) {
	type edge struct {
		to, rid, wt int
		sp          bool
	}
	g := make([][]edge, n)
	for _, e := range edges {
		x, y, wt := e[0], e[1], e[2]
		sp := wt < 0
		if sp {
			wt = 1 // 至少是 1
		}
		g[x] = append(g[x], edge{y, len(g[y]), wt, sp})
		g[y] = append(g[y], edge{x, len(g[x]) - 1, wt, sp})
	}

	var delta int
	dis := make([][2]int, n)
	for i := range dis {
		dis[i] = [2]int{1e18, 1e18}
	}
	dis[st] = [2]int{}
	dijkstra := func(k int) {
		h := hp{{st, 0}}
		for len(h) > 0 {
			p := h.pop()
			x := p.v
			if dis[x][k] < p.dis {
				continue
			}
			for i, e := range g[x] {
				y, wt := e.to, e.wt
				if k == 1 && e.sp {
					if newD := dis[y][0] + delta - dis[x][1]; newD > wt {
						wt = newD
						g[x][i].wt = newD
						g[y][e.rid].wt = newD
					}
				}
				if newD := dis[x][k] + wt; newD < dis[y][k] {
					dis[y][k] = newD
					h.push(pair{y, newD})
				}
			}
		}
	}

	dijkstra(0)
	delta = target - dis[end][0]
	if delta < 0 { // -1 全改为 1 时，最短路比 target 还大
		return
	}

	dijkstra(1)
	if dis[end][1] < target { // 最短路无法达到 target
		return
	}

	for v, es := range g {
		for _, e := range es {
			if e.to > v {
				ans = append(ans, []int{v, e.to, e.wt})
			}
		}
	}
	return
}

type pair struct{ v, dis int }
type hp []pair

func (h hp) Len() int              { return len(h) }
func (h hp) Less(i, j int) bool    { return h[i].dis < h[j].dis }
func (h hp) Swap(i, j int)         { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v interface{})   { *h = append(*h, v.(pair)) }
func (h *hp) Pop() (v interface{}) { a := *h; *h, v = a[:len(a)-1], a[len(a)-1]; return }
func (h *hp) push(v pair)          { heap.Push(h, v) }
func (h *hp) pop() pair            { return heap.Pop(h).(pair) }
