//go:build main
// +build main

package main

import (
	"bufio"
	"container/heap"
	. "fmt"
	"io"
	"os"
)

func CF938D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, m, from, to, wt int
	Fscan(in, &n, &m)
	const (
		inf int64 = 1e18
	)
	n++
	type edge struct{ to, wt int }
	g := make([][]edge, n)
	for i := 1; i <= m; i++ {
		Fscan(in, &from, &to, &wt)
		g[from] = append(g[from], edge{
			to: to,
			wt: wt * 2,
		})
		g[to] = append(g[to], edge{
			to: from,
			wt: wt * 2,
		})
	}
	for i := 1; i < n; i++ {
		Fscan(in, &wt)
		g[0] = append(g[0], edge{
			to: i,
			wt: wt,
		})
	}

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

	dis := dijkstra(g, 0)
	for i := 1; i < n; i++ {
		Fprint(out, dis[i], " ")
	}

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

func main() { CF938D(os.Stdin, os.Stdout) }
