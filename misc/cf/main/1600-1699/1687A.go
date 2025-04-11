//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF1687A(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, k int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &k)
		a := make([]int, n)
		s := 0
		for i := range a {
			Fscan(in, &a[i])
			if i < k {
				s += a[i]
			}
		}
		if k >= n {
			Fprintln(out, s+k*n-n*(n+1)/2)
			continue
		}
		ans := s
		for i, v := range a[k:] {
			s += v - a[i]
			ans = max(ans, s)
		}
		Fprintln(out, ans+k*(k-1)/2)
	}

}

func main() { CF1687A(os.Stdin, os.Stdout) }
