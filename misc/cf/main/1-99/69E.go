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

func CF69E(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, k int
	Fscan(in, &n, &k)
	a := make([]int, n)
	cnt := map[int]int{}
	h := hp{}
	for i := range a {
		Fscan(in, &a[i])
		v := a[i]
		cnt[v]++
		if cnt[v] == 1 {
			heap.Push(&h, v)
		}
		if i >= k-1 {
			for h.Len() > 0 && cnt[h.IntSlice[0]] != 1 {
				heap.Pop(&h)
			}
			if h.Len() > 0 {
				Fprintln(out, h.IntSlice[0])
			} else {
				Fprintln(out, "Nothing")
			}
			v := a[i-k+1]
			cnt[v]--
			if cnt[v] == 1 {
				heap.Push(&h, v)
			}
		}
	}
}

type hp struct{ sort.IntSlice }

func (h hp) Less(i, j int) bool { return h.IntSlice[i] > h.IntSlice[j] }
func (h *hp) Push(v any)        { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp) Pop() any          { a := h.IntSlice; v := a[len(a)-1]; h.IntSlice = a[:len(a)-1]; return v }

func main() { CF69E(os.Stdin, os.Stdout) }
