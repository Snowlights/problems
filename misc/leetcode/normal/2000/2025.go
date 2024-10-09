package _000

import "container/heap"

// 2039
func networkBecomesIdle(edges [][]int, patience []int) int {
	n := len(patience)

	const (
		inf int64 = 1e18
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
	for _, v := range edges {
		from, to := v[0], v[1]
		g[from] = append(g[from], edge{
			to: to,
			wt: 1,
		})
		g[to] = append(g[to], edge{
			to: from,
			wt: 1,
		})
	}

	dist := dijkstra(g, 0)
	// fmt.Println(dist)
	ans := 0
	for i := 1; i < n; i++ {
		res, onceCost := dist[i]*2, dist[i]*2
		if onceCost <= int64(patience[i]) {
			res = onceCost
		} else {
			last := onceCost % int64(patience[i])
			if last == 0 {
				last = onceCost - int64(patience[i])
			} else {
				last = onceCost - last
			}
			res = last + onceCost
		}
		ans = max(ans, int(res))
	}
	return ans + 1
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
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
