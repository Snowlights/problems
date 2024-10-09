//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"sort"
)

func CF650D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, m int
	Fscan(in, &n, &m)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	type pair struct{ qid, v int }
	qs := make([][]pair, n)
	for i := 0; i < m; i++ {
		var j, v int
		Fscan(in, &j, &v)
		qs[j-1] = append(qs[j-1], pair{i, v})
	}

	ans := make([]int, m)
	pre := make([]int, n)
	g := []int{}
	for i, v := range a {
		for _, p := range qs[i] {
			ans[p.qid] = sort.SearchInts(g, p.v) + 1
		}
		p := sort.SearchInts(g, v)
		if p < len(g) {
			g[p] = v
		} else {
			g = append(g, v)
		}
		pre[i] = p + 1
	}

	suf := make([]int, n)
	g = g[:0]
	for i := n - 1; i >= 0; i-- {
		for _, p := range qs[i] {
			ans[p.qid] += sort.SearchInts(g, -p.v)
		}
		v := -a[i]
		p := sort.SearchInts(g, v)
		if p < len(g) {
			g[p] = v
		} else {
			g = append(g, v)
		}
		suf[i] = p + 1
	}

	lis := len(g)
	cnt := make([]int, n+1)
	for i, p := range pre {
		if p+suf[i]-1 == lis {
			cnt[p]++
		}
	}

	for i, p := range pre {
		l := lis
		if p+suf[i]-1 == lis && cnt[p] == 1 {
			l--
		}
		for _, p := range qs[i] {
			ans[p.qid] = max(ans[p.qid], l)
		}
	}

	for _, v := range ans {
		Fprintln(out, v)
	}
}

func main() { CF650D(os.Stdin, os.Stdout) }
