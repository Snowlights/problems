//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF1714D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n int
	var t string
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &t, &n)
		a := make([]string, n)
		for i := range a {
			Fscan(in, &a[i])
		}

		n = len(t)
		f := make([]int, n+1)
		type pair struct{ ti, ai int }
		from := make([]pair, n+1)
		for i := 1; i <= n; i++ {
			f[i] = 1e9
			for j, s := range a {
				m := len(s)
				for k := max(i-m, 0); k < i && k <= n-m; k++ {
					if t[k:k+m] == s && f[k]+1 < f[i] {
						f[i] = f[k] + 1
						from[i] = pair{k, j + 1}
					}
				}
			}
		}

		if f[n] == 1e9 {
			Fprintln(out, -1)
			continue
		}
		Fprintln(out, f[n])
		i := n
		for ; from[i].ti > 0; i = from[i].ti {
			Fprintln(out, from[i].ai, from[i].ti+1)
		}
		Fprintln(out, from[i].ai, 1)
	}

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() { CF1714D(os.Stdin, os.Stdout) }
