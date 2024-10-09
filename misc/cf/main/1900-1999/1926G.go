//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF1926G(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	T, n, s := 0, 0, ""
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		p := make([]int, n+1)
		for i := 2; i <= n; i++ {
			Fscan(in, &p[i])
		}
		Fscan(in, &s)
		f := make([][2]int, n+1)
		for i := n; i > 0; i-- {
			g := f[i]
			if s[i-1] != 'C' {
				g[s[i-1]&1] = n
			}
			f[p[i]][0] += min(g[0], g[1]+1)
			f[p[i]][1] += min(g[1], g[0]+1)
		}
		Fprintln(out, min(f[0][0], f[0][1]))
	}

}

func main() { CF1926G(os.Stdin, os.Stdout) }
