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

func CF1106D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, m, v, w int
	Fscan(in, &n, &m)
	g := make([][]int, n+1)
	for ; m > 0; m-- {
		Fscan(in, &v, &w)
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}
	h := hp{[]int{1}}
	vis := make([]bool, n+1)
	vis[1] = true
	for h.Len() > 0 {
		v := heap.Pop(&h).(int)
		Fprint(out, v, " ")
		for _, w := range g[v] {
			if !vis[w] {
				vis[w] = true
				heap.Push(&h, w)
			}
		}
	}
}

type hp struct{ sort.IntSlice }

func (h *hp) Push(v interface{}) {
	h.IntSlice = append(h.IntSlice, v.(int))
}

func (h *hp) Pop() interface{} {
	a := h.IntSlice
	v := a[len(a)-1]
	h.IntSlice = a[:len(a)-1]
	return v
}

func main() { CF1106D(os.Stdin, os.Stdout) }
