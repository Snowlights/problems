//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF354A(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, l, r, ql, qr, pre, suf int
	Fscan(in, &n, &l, &r, &ql, &qr)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
		suf += a[i]
	}
	ans := suf*r + (n-1)*qr
	for i, v := range a {
		suf -= v
		pre += v
		cost := pre*l + suf*r
		cl, cr := i+1, n-1-i
		if cl < cr {
			cost += (cr - cl - 1) * qr
		} else if cr < cl {
			cost += (cl - cr - 1) * ql
		}
		ans = min(ans, cost)
	}
	Fprint(out, ans)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() { CF354A(os.Stdin, os.Stdout) }
