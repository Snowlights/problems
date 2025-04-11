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

func CF1974G(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, x, c int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &x)
		ans, s := n, 0
		h := &hp{}
		for ; n > 0; n-- {
			Fscan(in, &c)
			s -= c
			heap.Push(h, c)
			for s < 0 {
				s += heap.Pop(h).(int)
				ans--
			}
			s += x
		}
		Fprintln(out, ans)
	}

}

type hp struct{ sort.IntSlice }

func (h hp) Less(i, j int) bool { return h.IntSlice[i] > h.IntSlice[j] }
func (h *hp) Push(v any)        { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp) Pop() any          { a := h.IntSlice; v := a[len(a)-1]; h.IntSlice = a[:len(a)-1]; return v }

func main() { CF1974G(os.Stdin, os.Stdout) }
