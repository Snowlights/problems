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

func CF371E(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, k, s, ss, mnR int
	Fscan(in, &n)
	a := make([]struct{ v, i int }, n)
	for i := range a {
		Fscan(in, &a[i].v)
		a[i].i = i + 1
	}
	Fscan(in, &k)

	sort.Slice(a, func(i, j int) bool { return a[i].v < a[j].v })
	mn := int(9e18)
	for r, p := range a {
		ss += p.v*min(r-1, k-1) - s
		s += p.v
		if l := r + 1 - k; l >= 0 {
			if ss < mn {
				mn = ss
				mnR = r + 1
			}
			s -= a[l].v
			ss -= s - a[l].v*(k-1)
		}
	}
	for _, p := range a[mnR-k : mnR] {
		Fprint(out, p.i, " ")
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() { CF371E(os.Stdin, os.Stdout) }
