//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF1626C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a, b := make([]int64, n), make([]int64, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		for i := range b {
			Fscan(in, &b[i])
		}

		minMagic := func() int64 {
			var mm, p, c int64
			prev := a[n-1]
			for i := n - 1; i >= 0; i-- {
				k, h := a[i], b[i]
				c -= prev - k
				if c <= 0 {
					mm += (p * (p + 1)) / 2
					p, c = 0, 0
				}
				if h > c {
					p += h - c
					c += h - c
				}
				prev = a[i]
			}
			mm += (p * (p + 1)) / 2
			return mm
		}
		Fprintln(out, minMagic())
	}

}

func main() { CF1626C(os.Stdin, os.Stdout) }
