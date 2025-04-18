//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF1672H(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, q, l, r int
	var s string
	Fscan(in, &n, &q, &s)
	sum := make([][2]int, n)
	for i := range n - 1 {
		sum[i+1] = sum[i]
		if s[i] == s[i+1] {
			sum[i+1][s[i]-'0']++
		}
	}
	for range q {
		Fscan(in, &l, &r)
		c0 := sum[r-1][0] - sum[l-1][0]
		c1 := sum[r-1][1] - sum[l-1][1]
		Fprintln(out, max(c0, c1)+1)
	}

}

func main() { CF1672H(os.Stdin, os.Stdout) }
