package main

import (
	"bufio"
	"container/heap"
	. "fmt"
	"io"
	"os"
	"sort"
)

func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, k, l, r int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([][]int, n+1)
		b := make([][]int, n)
		ans := 0
		for i := 0; i < n; i++ {
			Fscan(in, &k, &l, &r)
			if l > r {
				ans += r
				a[k] = append(a[k], l-r)
			} else {
				ans += l
				b[n-k] = append(b[n-k], r-l)
			}
		}
		f := func(a [][]int) {
			h := &hp{}
			for k, ds := range a {
				for _, d := range ds {
					ans += d
					heap.Push(h, d)
					if h.Len() > k {
						ans -= heap.Pop(h).(int)
					}
				}
			}
		}
		f(a)
		f(b)
		Fprintln(out, ans)
	}

}

type hp struct{ sort.IntSlice }

func (h *hp) Push(v any) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp) Pop() any   { a := h.IntSlice; v := a[len(a)-1]; h.IntSlice = a[:len(a)-1]; return v }

func main() { run(os.Stdin, os.Stdout) }
