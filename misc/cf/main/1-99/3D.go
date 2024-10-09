//go:build main
// +build main

package main

import (
	"bufio"
	"container/heap"
	. "fmt"
	"io"
	"os"
)

func CF3D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var (
		s              []byte
		l, r, ans, cnt int
	)
	Fscan(in, &s)
	h := &hp{}
	for i, c := range s {
		if c == '(' {
			cnt++
			continue
		}
		if c == '?' {
			s[i] = ')'
			Fscan(in, &l, &r)
			ans += r
			heap.Push(h, pair{l - r, i})
		}
		if cnt > 0 {
			cnt--
			continue
		}
		if h.Len() == 0 {
			Fprintln(out, -1)
			return
		}
		cnt++
		p := heap.Pop(h).(pair)
		ans += p.v
		s[p.i] = '('
	}
	if cnt > 0 {
		Fprintln(out, -1)
		return
	}
	Fprintln(out, ans)
	Fprintln(out, string(s))

}

type pair struct{ v, i int }
type hp []pair

func (h hp) Len() int            { return len(h) }
func (h hp) Less(i, j int) bool  { return h[i].v < h[j].v }
func (h hp) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v interface{}) { *h = append(*h, v.(pair)) }
func (h *hp) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }

func main() { CF3D(os.Stdin, os.Stdout) }
