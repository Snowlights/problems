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

func CF1837F(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, k int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &k)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		suf := make([]int, n+1)
		ans := sort.Search(k*1e9, func(mx int) bool {
			h := hp{}
			s := mx
			upd := func(v int) {
				if s >= v {
					s -= v
					heap.Push(&h, v)
				} else if h.Len() > 0 && v < h.IntSlice[0] {
					s += h.IntSlice[0] - v
					h.IntSlice[0] = v
					heap.Fix(&h, 0)
				}
			}
			for i := n - 1; i >= 0; i-- {
				upd(a[i])
				suf[i] = h.Len()
			}
			if h.Len() >= k {
				return true
			}

			h = hp{}
			s = mx
			for i, v := range a {
				upd(v)
				if h.Len()+suf[i+1] >= k {
					return true
				}
			}
			return false
		})
		Fprintln(out, ans)
	}
}

type hp struct{ sort.IntSlice }

func (h hp) Less(i, j int) bool { return h.IntSlice[i] > h.IntSlice[j] }
func (h *hp) Push(v interface{})        { h.IntSlice = append(h.IntSlice, v.(int)) }
func (hp) Pop() (_ interface{})         { return }

func main() { CF1837F(os.Stdin, os.Stdout) }
