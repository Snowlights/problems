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

func CF1862E(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, m, d, v int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m, &d)
		ans, s := 0, 0
		h := hp{}
		for i := 1; i <= n; i++ {
			Fscan(in, &v)
			if v <= 0 {
				continue
			}
			if h.Len() < m {
				s += v
				heap.Push(&h, v)
			} else if v > h.IntSlice[0] {
				s += v - h.replace(v)
			}
			ans = max(ans, s-i*d)
		}
		Fprintln(out, ans)
	}
}

type hp struct{ sort.IntSlice }

func (h *hp) Push(v any)        { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp) Pop() any          { a := h.IntSlice; v := a[len(a)-1]; h.IntSlice = a[:len(a)-1]; return v }
func (h *hp) replace(v int) int { top := h.IntSlice[0]; h.IntSlice[0] = v; heap.Fix(h, 0); return top }

func main() { CF1862E(os.Stdin, os.Stdout) }
