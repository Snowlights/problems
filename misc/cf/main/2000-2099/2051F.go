//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF2051F(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, m, q, v int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m, &q)
		h, t, l, r := 0, n+1, m, m
		for range q {
			Fscan(in, &v)
			if h > 0 && v > h {
				h++
			}
			if t <= n && v < t {
				t--
			}
			if v < l {
				l--
			} else if v > r {
				r++
			} else {
				if l == r {
					l = -1
				}
				h = max(h, 1)
				t = min(t, n)
			}
			if l < 0 {
				Fprint(out, min(h+n+1-t, n), " ")
			} else {
				Fprint(out, r-l+1+min(h, l-1)+n+1-max(t, r+1), " ")
			}
		}
		Fprintln(out)
	}

}

func main() { CF2051F(os.Stdin, os.Stdout) }
