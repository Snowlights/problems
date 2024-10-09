//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF633F(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, ans, v, w int
	Fscan(in, &n)
	a := make([]int, n+1)
	for i := 1; i <= n; i++ {
		Fscan(in, &a[i])
	}
	g := make([][]int, n+1)
	for i := 1; i < n; i++ {
		Fscan(in, &v, &w)
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}
	var F func(int, int) (int, int, int)
	F = func(x, f int) (int, int, int) {
		v := a[x]
		C, P2, P, CP := v, v, 0, v
		for _, y := range g[x] {
			if y == f {
				continue
			}
			c, p, cp := F(y, x)
			ans = max(ans, P2+p, P+p, C+cp, CP+c)
			CP = max(CP, cp+v, P+c+v, C+p)
			P2 = max(P2, C+c)
			P = max(P, p)
			C = max(C, c+v)
		}
		return C, max(P2, P), CP
	}
	F(1, 0)
	Fprintln(out, ans)

}

func main() { CF633F(os.Stdin, os.Stdout) }
