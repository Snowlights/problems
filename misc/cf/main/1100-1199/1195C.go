//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF1195C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n int
	Fscan(in, &n)
	a := make([][]int, 2)
	for i := range a {
		a[i] = make([]int, n)
		for j := range a[i] {
			Fscan(in, &a[i][j])
		}
	}
	f0, f1, f2 := 0, a[0][0], a[1][0]
	for i := 1; i < n; i++ {
		f0, f1, f2 = max(f0, max(f1, f2)), max(f0, f2)+a[0][i], max(f0, f1)+a[1][i]
	}

	Fprintln(out, max(f0, max(f1, f2)))
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func main() { CF1195C(os.Stdin, os.Stdout) }
