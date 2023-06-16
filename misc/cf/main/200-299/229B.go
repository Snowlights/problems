//go:build main
// +build main

package main

import (
	"bufio"
	"container/heap"
	. "fmt"
	"io"
	"os"
	"sort"
)

func CF229B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	const (
		inf int64 = 1e18
	)
	var n, m, v, w, wt int

	type edge struct{ to, wt int }
	Fscan(in, &n, &m)
	g := make([][]edge, n)
	for ; m > 0; m-- {
		Fscan(in, &v, &w, &wt)
		v--
		w--
		g[v] = append(g[v], edge{w, wt})
		g[w] = append(g[w], edge{v, wt})
	}
	type seg struct{ l, r int }
	t := make([][]seg, n)
	for i := 0; i < n-1; i++ {
		pre := -2
		for Fscan(in, &m); m > 0; m-- {
			Fscan(in, &v)
			if v > pre+1 {
				t[i] = append(t[i], seg{v, v})
			} else {
				t[i][len(t[i])-1].r = v
			}
			pre = v
		}
	}

	dijkstra := func(g [][]edge, start int) []int64 {
		dist := make([]int64, n)
		for i := range dist {
			dist[i] = inf
		}
		dist[start] = 0
		if len(t[0]) > 0 && t[0][0].l == 0 {
			dist[0] = int64(t[0][0].r + 1)
		}
		// 虽然可以用 dist 来判断是否需要 relax，但是对于一些变形题，用 vis 是最稳的
		vis := make([]bool, n)
		fa := make([]int, n)
		for i := range fa {
			fa[i] = -1
		}
		h := vdHeap{{start, dist[0]}}
		for len(h) > 0 {
			p := h.pop()
			v := p.v
			if vis[v] { // p.dis > dist[v]
				continue
			}
			vis[v] = true
			for _, e := range g[v] {
				w, wt := e.to, e.wt
				newD := dist[v] + int64(wt)
				t := t[w]
				i := sort.Search(len(t), func(i int) bool { return int64(t[i].r) >= newD })

				if i < len(t) && newD >= int64(t[i].l) {
					newD = int64(t[i].r + 1)
				}

				if newD < dist[w] {
					dist[w] = newD
					fa[w] = v
					h.push(vdPair{w, dist[w]})
				}
			}
		}
		return dist
	}

	dist := dijkstra(g, 0)
	if dist[n-1] == inf {
		Fprintln(out, -1)
	} else {
		Fprintln(out, dist[n-1])
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

func main() { CF229B(os.Stdin, os.Stdout) }
