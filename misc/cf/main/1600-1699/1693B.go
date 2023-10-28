//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
)

func CF1693B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, T, p int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		g := make([][]int, n)
		for i := 1; i < n; i++ {
			Fscan(in, &p)
			g[p-1] = append(g[p-1], i)
		}
		lim := make([]struct{ l, r int64 }, n)
		for i := range lim {
			Fscan(in, &lim[i].l, &lim[i].r)
		}
		ans := 0
		var f func(int) int64
		f = func(v int) (sum int64) {
			for _, w := range g[v] {
				sum += f(w)
			}
			if sum < lim[v].l {
				ans++
				sum = lim[v].r
			} else if sum > lim[v].r {
				sum = lim[v].r
			}
			return
		}
		f(0)
		Fprintln(out, ans)
	}

}

// func main() { CF1693B(os.Stdin, os.Stdout) }
