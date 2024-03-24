//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF777C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, m, q, l, r int
	Fscan(in, &n, &m)
	minL, up := make([]int, n), make([]int, m)
	a := make([][]int, n)
	for i := range a {
		minL[i] = i
		a[i] = make([]int, m)
		for j := range a[i] {
			Fscan(in, &a[i][j])
			if i > 0 && a[i][j] < a[i-1][j] {
				up[j] = i
			}
			minL[i] = min(minL[i], up[j])
		}
	}
	for Fscan(in, &q); q > 0; q-- {
		Fscan(in, &l, &r)
		if minL[r-1] <= l-1 {
			Fprintln(out, "Yes")
		} else {
			Fprintln(out, "No")
		}
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() { CF777C(os.Stdin, os.Stdout) }
