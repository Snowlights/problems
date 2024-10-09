//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF1772D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, pre, v int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &pre)
		l, r := 0, int(1e9)
		for ; n > 1; n-- {
			Fscan(in, &v)
			if pre > v {
				l = max(l, (pre+v+1)/2)
			} else if pre < v {
				r = min(r, (pre+v)/2)
			}
			pre = v
		}
		if l <= r {
			Fprintln(out, l)
		} else {
			Fprintln(out, -1)
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
	if a > b {
		return b
	}
	return a
}

func main() { CF1772D(os.Stdin, os.Stdout) }
