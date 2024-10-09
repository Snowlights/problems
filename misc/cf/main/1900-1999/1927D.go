//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF1927D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, pre, v, q, l, r int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &pre)
		left := make([]int, n+1)
		for i := 2; i <= n; i++ {
			Fscan(in, &v)
			if v == pre {
				left[i] = left[i-1]
			} else {
				left[i] = i - 1
			}
			pre = v
		}
		for Fscan(in, &q); q > 0; q-- {
			Fscan(in, &l, &r)
			if left[r] < l {
				Fprintln(out, -1, -1)
			} else {
				Fprintln(out, left[r], r)
			}
		}
		Fprintln(out)
	}

}

func main() { CF1927D(os.Stdin, os.Stdout) }
