package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, m, k, w, v, c int
	Fscan(in, &n, &m, &k)
	type pair struct{ w, v int }
	g := make([][]pair, n)
	for i := 0; i < n; i++ {
		Fscan(in, &w, &v, &c)
		g[c-1] = append(g[c-1], pair{w, v})
	}

	f := make([]int, m+1)
	for _, a := range g {
		nf := append(f[:0:0], f...)
		for _, p := range a {
			for j := m; j >= p.w; j-- {
				nf[j] = max(nf[j], f[j-p.w]+p.v+k, nf[j-p.w]+p.v)
			}
		}
		f = nf
	}
	Fprint(out, f[m])

}

func main() { run(os.Stdin, os.Stdout) }
