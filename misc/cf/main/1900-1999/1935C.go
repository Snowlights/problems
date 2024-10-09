//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF1935C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, lim int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &lim)
		type pair struct{ a, b int }
		a := make([]pair, n)
		for i := range a {
			Fscan(in, &a[i].a, &a[i].b)
		}
		slices.SortFunc(a, func(p, q pair) int { return p.b - q.b })

		ans := 0
		f := make([]int, n)
		for j, p := range a {
			f[j] = p.a - p.b
			if p.a <= lim {
				ans = 1
			}
		}
		for i := 1; i < n; i++ {
			mn := f[i-1]
			for j := i; j < n; j++ {
				f[j], mn = mn+a[j].a, min(mn, f[j])
				if f[j]+a[j].b <= lim {
					ans = i + 1
				}
			}
		}
		Fprintln(out, ans)
	}

}

func main() { CF1935C(os.Stdin, os.Stdout) }
