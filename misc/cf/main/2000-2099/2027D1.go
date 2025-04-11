//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF2027D1(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, m int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m)
		s := make([]int, n+1)
		for i := 1; i <= n; i++ {
			Fscan(in, &s[i])
			s[i] += s[i-1]
		}
		b := make([]int, m)
		for i := range b {
			Fscan(in, &b[i])
		}
		f := make([]int, n+1)
		for j := range n {
			f[j] = 1e18
		}
		for i := m - 1; i >= 0; i-- {
			k := n
			for j := n - 1; j >= 0; j-- {
				for s[k] > s[j]+b[i] {
					k--
				}
				f[j] = min(f[j], f[k]+m-1-i)
			}
		}
		ans := f[0]
		if ans >= 1e18 {
			ans = -1
		}
		Fprintln(out, ans)
	}

}

func main() { CF2027D1(os.Stdin, os.Stdout) }
