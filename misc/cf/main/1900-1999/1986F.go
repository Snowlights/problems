//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF1986F(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, m, v, w, clock int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m)
		g := make([][]int, n)
		for range m {
			Fscan(in, &v, &w)
			v--
			w--
			g[v] = append(g[v], w)
			g[w] = append(g[w], v)
		}
		remove := 0
		dfn := make([]int, n)
		var tarjan func(int, int) (int, int)
		tarjan = func(v, fa int) (int, int) {
			clock++
			dfn[v] = clock
			lowV := clock
			size := 1
			for _, w := range g[v] {
				if dfn[w] == 0 {
					lowW, sz := tarjan(w, v)
					lowV = min(lowV, lowW)
					if lowW > dfn[v] {
						remove = max(remove, sz*(n-sz))
					}
					size += sz
				} else if w != fa {
					lowV = min(lowV, dfn[w])
				}
			}
			return lowV, size
		}
		for v, ts := range dfn {
			if ts == 0 {
				tarjan(v, -1)
			}
		}
		Fprintln(out, n*(n-1)/2-remove)
	}

}

func main() { CF1986F(os.Stdin, os.Stdout) }
