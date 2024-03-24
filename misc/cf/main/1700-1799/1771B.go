//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF1771B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, m, x, y int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m)
		l := make([]int, n+1)
		for ; m > 0; m-- {
			Fscan(in, &x, &y)
			if x > y {
				x, y = y, x
			}
			l[y] = max(l[y], x)
		}
		ans, maxL := 0, 0
		for i, x := range l {
			maxL = max(maxL, x)
			ans += i - maxL
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

func main() { CF1771B(os.Stdin, os.Stdout) }
