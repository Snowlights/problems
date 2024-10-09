//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF1873F(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, k int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &k)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		h := make([]int, n)
		for i := range h {
			Fscan(in, &h[i])
		}
		ans := 0
		for i := 0; i < n; {
			st := i
			for i++; i < n && h[i-1]%h[i] == 0; i++ {
			}
			b := a[st:i]
			s, l := 0, 0
			for r, v := range b {
				s += v
				for s > k {
					s -= b[l]
					l++
				}
				ans = max(ans, r-l+1)
			}
		}
		Fprintln(out, ans)
	}

}

func main() { CF1873F(os.Stdin, os.Stdout) }
