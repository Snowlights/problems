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

func CF1140C(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, k, ans, sum int
	Fscan(in, &n, &k)
	a := make([]struct{ len, b int }, n)
	for i := range a {
		Fscan(in, &a[i].len, &a[i].b)
	}
	sort.Slice(a, func(i, j int) bool { return a[i].b > a[j].b })

	h := &hp{}
	for _, p := range a {
		heap.Push(h, p.len)
		sum += p.len
		for h.Len() > k {
			sum -= heap.Pop(h).(int)
		}
		ans = max(ans, sum*p.b)
	}
	Fprint(out, ans)
}

type hp struct{ sort.IntSlice }

func (h *hp) Push(v interface{}) { h.IntSlice = append(h.IntSlice, v.(int)) }

func (h *hp) Pop() interface{} {
	a := h.IntSlice
	v := a[len(a)-1]
	h.IntSlice = a[:len(a)-1]
	return v
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() { CF1140C(os.Stdin, os.Stdout) }
