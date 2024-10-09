package _900_2000

import "container/heap"

// 1976
func countPaths(n int, roads [][]int) int {

	const (
		inf int64 = 1e18
		mod int   = 1e9 + 7
	)

	type edge struct{ to, wt int }
	dijkstra := func(g [][]edge, start int) []int64 {
		dist := make([]int64, n)
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
				if newD := dist[v] + int64(e.wt); newD < dist[w] {
					dist[w] = newD
					fa[w] = v
					h.push(vdPair{w, dist[w]})
				}
			}
		}
		return dist
	}

	g := make([][]edge, n)
	for _, v := range roads {
		from, to, wt := v[0], v[1], v[2]
		g[from] = append(g[from], edge{
			to: to,
			wt: wt,
		})
		g[to] = append(g[to], edge{
			to: from,
			wt: wt,
		})
	}

	dist := dijkstra(g, 0)
	indegree := make(map[int]int)
	for from, to := range g {
		for _, wt := range to {
			if dist[from]+int64(wt.wt) == dist[wt.to] {
				indegree[wt.to]++
			}
		}
	}

	q, dp := []int{0}, make([]int, n)
	dp[0] = 1
	for len(q) > 0 {
		val := q[0]
		q = q[1:]

		for _, v := range g[val] {
			if dist[val]+int64(v.wt) == dist[v.to] {
				dp[v.to] += dp[val]
				dp[v.to] %= mod

				if indegree[v.to]--; indegree[v.to] == 0 {
					q = append(q, v.to)
				}
			}
		}
	}

	return dp[n-1]
}

type vdPair struct {
	v   int
	dis int64
}
type vdHeap []vdPair

func (h vdHeap) Len() int              { return len(h) }
func (h vdHeap) Less(i, j int) bool    { return h[i].dis < h[j].dis }
func (h vdHeap) Swap(i, j int)         { h[i], h[j] = h[j], h[i] }
func (h *vdHeap) Push(v interface{})   { *h = append(*h, v.(vdPair)) }
func (h *vdHeap) Pop() (v interface{}) { a := *h; *h, v = a[:len(a)-1], a[len(a)-1]; return }
func (h *vdHeap) push(v vdPair)        { heap.Push(h, v) }
func (h *vdHeap) pop() vdPair          { return heap.Pop(h).(vdPair) }
