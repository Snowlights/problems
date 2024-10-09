//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF1304C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var q, n, t, l, r int
	var ok bool
	for Fscan(in, &q); q > 0; q-- {
		Fscan(in, &n, &t)
		ok = true
		for ll, rr, pre := t, t, 0; n > 0; n-- {
			Fscan(in, &t, &l, &r)
			ll, rr = ll-t+pre, rr+t-pre
			if l > rr || r < ll {
				ok = false
			}
			ll = max(ll, l)
			rr = min(rr, r)
			pre = t
		}
		if ok {
			Fprintln(out, "YES")
		} else {
			Fprintln(out, "NO")
		}
	}

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() { CF1304C(os.Stdin, os.Stdout) }
