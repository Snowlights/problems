package main

import (
	"container/heap"
)

func minimumTime(n int, edges [][]int, disappear []int) []int {
	const (
		inf int = 1e9
	)
	type edge struct{ to, wt int }
	g := make([][]edge, n)
	for _, v := range edges {
		g[v[0]] = append(g[v[0]], edge{
			to: v[1],
			wt: v[2],
		})
		g[v[1]] = append(g[v[1]], edge{
			to: v[0],
			wt: v[2],
		})
	}
	dijkstra := func(g [][]edge, start int) []int {
		dist := make([]int, n)
		for i := range dist {
			dist[i] = inf
		}
		dist[start] = 0
		// 虽然可以用 dist 来判断是否需要 relax，但是对于一些变形题，用 vis 是最稳的
		vis := make([]bool, n)
		fa := make([]int, n)
		for i := range fa {
			fa[i] = -1
		}
		h := vdHeap{{start, 0}}
		for len(h) > 0 {
			p := h.pop()
			v := p.v
			if vis[v] { // p.dis > dist[v]
				continue
			}
			vis[v] = true
			for _, e := range g[v] {
				w := e.to
				if newD := dist[v] + int(e.wt); newD < disappear[w] && newD < dist[w] {
					dist[w] = newD
					fa[w] = v
					h.push(vdPair{w, dist[w]})
				}
			}
		}
		return dist
	}
	dis := dijkstra(g, 0)
	ans := make([]int, n)
	for i, v := range disappear {
		if dis[i] < v {
			ans[i] = dis[i]
		} else {
			ans[i] = -1
		}
	}

	return ans
}

type vdPair struct {
	v   int
	dis int
}
type vdHeap []vdPair

func (h vdHeap) Len() int              { return len(h) }
func (h vdHeap) Less(i, j int) bool    { return h[i].dis < h[j].dis }
func (h vdHeap) Swap(i, j int)         { h[i], h[j] = h[j], h[i] }
func (h *vdHeap) Push(v interface{})   { *h = append(*h, v.(vdPair)) }
func (h *vdHeap) Pop() (v interface{}) { a := *h; *h, v = a[:len(a)-1], a[len(a)-1]; return }
func (h *vdHeap) push(v vdPair)        { heap.Push(h, v) }
func (h *vdHeap) pop() vdPair          { return heap.Pop(h).(vdPair) }
