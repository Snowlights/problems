//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF1969D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, k int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &k)
		type pair struct{ x, y int }
		a := make([]pair, n)
		for i := range a {
			Fscan(in, &a[i].x)
		}
		for i := range a {
			Fscan(in, &a[i].y)
		}
		slices.SortFunc(a, func(a, b pair) int { return b.y - a.y })

		s := 0
		h := hp{make([]int, k)}
		for i, p := range a[:k] {
			h.IntSlice[i] = p.x
			s -= p.x
		}
		heap.Init(&h)

		for _, p := range a[k:] {
			s += max(p.y-p.x, 0)
		}
		ans := max(s, 0)
		for _, p := range a[k:] {
			s -= max(p.y-p.x, 0)
			if k > 0 && p.x < h.IntSlice[0] {
				s += h.IntSlice[0] - p.x
				ans = max(ans, s)
				h.IntSlice[0] = p.x
				heap.Fix(&h, 0)
			}
		}
		Fprintln(out, ans)
	}

}

type hp struct{ sort.IntSlice }
func (h hp) Less(i, j int) bool { return h.IntSlice[i] > h.IntSlice[j] }
func (hp) Push(any)             {}
func (hp) Pop() (_ any)         { return }

func main() { CF1969D(os.Stdin, os.Stdout) }
