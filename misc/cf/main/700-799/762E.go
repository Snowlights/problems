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

func CF762E(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, k, ans int
	Fscan(in, &n, &k)
	type tuple struct{ x, r, f int }
	a := make([]tuple, n)
	g := map[int][]int{}
	for i := range a {
		Fscan(in, &a[i].x, &a[i].r, &a[i].f)
		g[a[i].f] = append(g[a[i].f], a[i].x)
	}
	slices.SortFunc(a, func(a, b tuple) int { return b.r - a.r })

	tree := map[int][]int{}
	for f, xs := range g {
		slices.Sort(xs)
		tree[f] = make([]int, len(xs)+1)
	}
	add := func(f, i int) {
		t := tree[f]
		for i = sort.SearchInts(g[f], i) + 1; i < len(t); i += i & -i {
			t[i]++
		}
	}
	pre := func(f, i int) (res int) {
		t := tree[f]
		for i = sort.SearchInts(g[f], i); i > 0; i &= i - 1 {
			res += t[i]
		}
		return
	}

	for _, p := range a {
		for f := p.f - k; f <= p.f+k; f++ {
			ans += pre(f, p.x+p.r+1) - pre(f, p.x-p.r)
		}
		add(p.f, p.x)
	}
	Fprint(out, ans)

}

func main() { CF762E(os.Stdin, os.Stdout) }
