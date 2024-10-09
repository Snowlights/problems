//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF1744F(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, v int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		pos := make([]int, n)
		for i := 0; i < n; i++ {
			Fscan(in, &v)
			pos[v] = i
		}

		ans := 1
		l, r := pos[0], pos[0]
		for m := 2; m <= n; m++ {
			i := pos[(m-1)/2]
			l = min(l, min(i, n-m))
			r = max(r, max(i, m-1))
			ans += max(m-(r-l), 0)
		}
		Fprintln(out, ans)
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

func main() { CF1744F(os.Stdin, os.Stdout) }
