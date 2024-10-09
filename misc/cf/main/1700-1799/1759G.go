//go:build main
// +build main

package main

import (
	"bufio"
	"container/heap"
	. "fmt"
	"io"
	"sort"
)

func CF1759G(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n int
o:
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a, pos := make([]int, n), make([]int, n+1)
		for i := 1; i < n; i += 2 {
			Fscan(in, &a[i])
			pos[a[i]] = i
		}
		hp := &hp{}
		for i := n; i > 0; i-- {
			if pos[i] > 0 {
				heap.Push(hp, pos[i])
			} else if hp.Len() > 0 {
				a[heap.Pop(hp).(int)-1] = i
			} else {
				Fprintln(out, -1)
				continue o
			}
		}
		for _, v := range a {
			Fprint(out, v, " ")
		}
		Fprintln(out)
	}

}

// func main() { CF1759G(os.Stdin, os.Stdout) }

type hp struct{ sort.IntSlice }

func (h hp) Less(i, j int) bool  { return h.IntSlice[i] > h.IntSlice[j] }
func (h *hp) Push(v interface{}) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp) Pop() interface{} {
	a := h.IntSlice
	v := a[len(a)-1]
	h.IntSlice = a[:len(a)-1]
	return v
}
