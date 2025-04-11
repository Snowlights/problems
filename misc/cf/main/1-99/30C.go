//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF30C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n int
	Fscan(in, &n)
	type tuple struct {
		x, y, t int
		p       float64
	}
	a := make([]tuple, n)
	for i := range a {
		Fscan(in, &a[i].x, &a[i].y, &a[i].t, &a[i].p)
	}
	slices.SortFunc(a, func(a, b tuple) int { return a.t - b.t })

	f := make([]float64, n)
	for i, t := range a {
		for j, q := range a[:i] {
			if (q.x-t.x)*(q.x-t.x)+(q.y-t.y)*(q.y-t.y) <= (q.t-t.t)*(q.t-t.t) {
				f[i] = max(f[i], f[j])
			}
		}
		f[i] += t.p
	}
	Fprintf(out, "%.10f", slices.Max(f))

}

func main() { CF30C(os.Stdin, os.Stdout) }
