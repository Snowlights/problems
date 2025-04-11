//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"slices"
)

func CF1986E(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, k, v int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &k)
		g := map[int][]int{}
		for i := 0; i < n; i++ {
			Fscan(in, &v)
			g[v%k] = append(g[v%k], v/k)
		}

		ans := 0
		odd := false
		for _, a := range g {
			slices.Sort(a)
			m := len(a)
			s := 0
			for i := m - 2; i >= 0; i -= 2 {
				s += a[i+1] - a[i]
			}
			if m%2 == 0 {
				ans += s
				continue
			}

			if odd {
				ans = -1
				break
			}
			odd = true

			minS := s
			for i := 1; i < m; i += 2 {
				s += a[i] - a[i-1] - (a[i+1] - a[i])
				minS = min(minS, s)
			}
			ans += minS
		}
		Fprintln(out, ans)
	}

}

func main() { CF1986E(os.Stdin, os.Stdout) }
