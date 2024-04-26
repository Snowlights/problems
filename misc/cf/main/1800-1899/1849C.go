//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF1849C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, m, l, r int
	var s string
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m, &s)
		ll, rr := make([]int, n+1), make([]int, n+2)
		for i, b := range s {
			if b == '0' {
				ll[i+1] = i + 1
			} else {
				ll[i+1] = ll[i]
			}
		}
		rr[n+1] = n + 1
		for i := n; i > 0; i-- {
			if s[i-1] == '1' {
				rr[i] = i
			} else {
				rr[i] = rr[i+1]
			}
		}

		set := map[[2]int]struct{}{}
		for ; m > 0; m-- {
			Fscan(in, &l, &r)
			l = rr[l]
			r = ll[r]
			if l > r {
				l, r = 0, 0
			}
			set[[2]int{l, r}] = struct{}{}
		}
		Fprintln(out, len(set))
	}

}

func main() { CF1849C(os.Stdin, os.Stdout) }
