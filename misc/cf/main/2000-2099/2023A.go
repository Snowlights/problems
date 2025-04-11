//go:build main
// +build main

package main

import (
	"bufio"
	"cmp"
	. "fmt"
	"io"
	"os"
	"slices"
)

func CF2023A(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		type pair struct{ x, y int }
		a := make([]pair, n)
		for i := range a {
			Fscan(in, &a[i].x, &a[i].y)
		}
		slices.SortFunc(a, func(a, b pair) int {
			return cmp.Or(min(a.x, a.y)-min(b.x, b.y), max(a.x, a.y)-max(b.x, b.y))
		})
		for _, p := range a {
			Fprint(out, p.x, p.y, " ")
		}
		Fprintln(out)
	}

}

func main() { CF2023A(os.Stdin, os.Stdout) }
