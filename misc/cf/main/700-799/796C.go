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

func CF796C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, v, w int
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	g := make([][]int, n)
	for i := 1; i < n; i++ {
		Fscan(in, &v, &w)
		v--
		w--
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}

	ans := int(2e9)
	h := hp{append(a[:0:0], a...)}
	heap.Init(&h)
	lazy := map[int]int{}
	for i, nb := range g {
		mx := a[i]
		lazy[mx]++
		for _, v := range nb {
			mx = max(mx, a[v]+1)
			lazy[a[v]]++
		}
		for h.Len() > 0 && lazy[h.IntSlice[0]] > 0 {
			lazy[h.IntSlice[0]]--
			heap.Pop(&h)
		}
		if h.Len() > 0 {
			mx = max(mx, h.IntSlice[0]+2)
		}
		ans = min(ans, mx)
		heap.Push(&h, a[i])
		for _, v := range nb {
			heap.Push(&h, a[v])
		}
	}
	Fprint(out, ans)
}

func main() { CF796C(os.Stdin, os.Stdout) }

type hp struct{ sort.IntSlice }

func (h hp) Less(i, j int) bool  { return h.IntSlice[i] > h.IntSlice[j] }
func (h *hp) Push(v interface{}) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp) Pop() interface{} {
	a := h.IntSlice
	v := a[len(a)-1]
	h.IntSlice = a[:len(a)-1]
	return v
}
func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}
func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}
