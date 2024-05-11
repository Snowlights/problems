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

func CF730I(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, p, s, ans int
	Fscan(in, &n, &p, &s)
	oriA := make([]int, n)
	a := make(hp, n)
	for i := range a {
		Fscan(in, &oriA[i])
		a[i] = pair{oriA[i], i}
	}
	sort.Slice(a, func(i, j int) bool { return a[i].v > a[j].v })

	group := make([]byte, n)
	for _, p := range a[:p] {
		ans += p.v
		group[p.i] = 1
	}
	a = a[p:]
	heap.Init(&a)

	oriB := make([]int, n)
	b := make(hp, 0, n-p)
	diff := make(hp, 0, p)
	for i, v := range oriA {
		Fscan(in, &oriB[i])
		if group[i] == 0 {
			b = append(b, pair{oriB[i], i})
		} else {
			diff = append(diff, pair{oriB[i] - v, i})
		}
	}
	heap.Init(&b)
	heap.Init(&diff)

	for ; s > 0; s-- {
		for group[a[0].i] > 0 {
			heap.Pop(&a)
		}
		for group[b[0].i] > 0 {
			heap.Pop(&b)
		}
		if b[0].v > a[0].v+diff[0].v { // 直接选 b
			ans += b[0].v
			group[b[0].i] = 2
			heap.Pop(&b)
		} else { // 反悔一个 a 变 b（a 那边选一个更小的）
			ans += a[0].v + diff[0].v
			group[a[0].i] = 1
			group[diff[0].i] = 2
			diff[0] = pair{oriB[a[0].i] - a[0].v, a[0].i}
			heap.Fix(&diff, 0)
			heap.Pop(&a)
		}
	}

	Fprintln(out, ans)
	for i, g := range group {
		if g == 1 {
			Fprint(out, i+1, " ")
		}
	}
	Fprintln(out)
	for i, g := range group {
		if g == 2 {
			Fprint(out, i+1, " ")
		}
	}

}

type pair struct{ v, i int }
type hp []pair

func (h hp) Len() int           { return len(h) }
func (h hp) Less(i, j int) bool { return h[i].v > h[j].v }
func (h hp) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v any)        { *h = append(*h, v.(pair)) }
func (h *hp) Pop() any          { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }

func main() { CF730I(os.Stdin, os.Stdout) }
