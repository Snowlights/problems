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

func CF923B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, t, s, ans int
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	hp := &hp2{}

	for _, v := range a {
		Fscan(in, &t)
		heap.Push(hp, v+s)
		ans = 0
		for hp.Len() > 0 && hp.IntSlice[0] <= s+t {
			ans += heap.Pop(hp).(int) - s
		}
		ans += hp.Len() * t
		Fprint(out, ans, " ")
		s += t
	}

}

type hp2 struct{ sort.IntSlice }

func (h hp2) Less(i, j int) bool  { return h.IntSlice[i] < h.IntSlice[j] }
func (h *hp2) Push(v interface{}) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp2) Pop() interface{} {
	a := h.IntSlice
	v := a[len(a)-1]
	h.IntSlice = a[:len(a)-1]
	return v
}

func main() { CF923B(os.Stdin, os.Stdout) }
