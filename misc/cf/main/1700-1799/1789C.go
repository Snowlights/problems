//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF1789C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, m, p, v int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m)
		a := make([]int, n)
		sumLen := make([]int, n+m+1)
		for i := range a {
			Fscan(in, &a[i])
			sumLen[a[i]] = m + 1
		}
		for i := 1; i <= m; i++ {
			Fscan(in, &p, &v)
			p--
			sumLen[a[p]] -= m + 1 - i
			a[p] = v
			sumLen[v] += m + 1 - i
		}
		ans := n * m * (m + 1)
		for _, k := range sumLen {
			ans -= k * (k - 1) / 2
		}
		Fprintln(out, ans)
	}

}

func main() { CF1789C(os.Stdin, os.Stdout) }
