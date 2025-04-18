//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"slices"
	"sort"
)

func CF1045G(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, k, ans int
	Fscan(in, &n, &k)
	type tuple struct{ x, r, iq int }
	a := make([]tuple, n)
	g := map[int][]int{}
	for i := range a {
		Fscan(in, &a[i].x, &a[i].r, &a[i].iq)
		g[a[i].iq] = append(g[a[i].iq], a[i].x)
	}
	slices.SortFunc(a, func(a, b tuple) int { return b.r - a.r })

	tree := map[int][]int{}
	for iq, xs := range g {
		slices.Sort(xs)
		tree[iq] = make([]int, len(xs)+1)
	}
	add := func(iq, i int) {
		t := tree[iq]
		for i = sort.SearchInts(g[iq], i) + 1; i < len(t); i += i & -i {
			t[i]++
		}
	}
	pre := func(iq, i int) (res int) {
		t := tree[iq]
		for i = sort.SearchInts(g[iq], i); i > 0; i &= i - 1 {
			res += t[i]
		}
		return
	}

	for _, p := range a {
		for iq := p.iq - k; iq <= p.iq+k; iq++ {
			ans += pre(iq, p.x+p.r+1) - pre(iq, p.x-p.r)
		}
		add(p.iq, p.x)
	}
	Fprint(out, ans)

}

func main() { CF1045G(os.Stdin, os.Stdout) }
