//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF1343C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, ans, mx int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		ans = 0
		for i := 0; i < n; {
			mx = a[i]
			for i++; i < n && a[i] > 0 == (mx > 0); i++ {
				mx = max(a[i], mx)
			}
			ans += mx
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

func main() { CF1343C(os.Stdin, os.Stdout) }
