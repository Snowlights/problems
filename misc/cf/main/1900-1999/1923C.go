//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF1923C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, q, v, l, r int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &q)
		s := make([]int, n+1)
		for i := range n {
			Fscan(in, &v)
			if v == 1 {
				s[i+1] = s[i] - 1
			} else {
				s[i+1] = s[i] + v - 1
			}
		}
		for range q {
			Fscan(in, &l, &r)
			if l < r && s[l-1] <= s[r] {
				Fprintln(out, "YES")
			} else {
				Fprintln(out, "NO")
			}
		}
	}

}

func main() { CF1923C(os.Stdin, os.Stdout) }
