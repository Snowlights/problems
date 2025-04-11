//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF2026B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		if n == 1 {
			Fprintln(out, 1)
			continue
		}
		suf := make([]int, n/2+1)
		for i := n - 2; i >= 0; i -= 2 {
			suf[i/2] = max(suf[i/2+1], a[i+1]-a[i])
		}
		if n%2 == 0 {
			Fprintln(out, suf[0])
			continue
		}
		ans, pre := suf[0], 0
		for i := 1; i < n; i += 2 {
			pre = max(pre, a[i]-a[i-1])
			ans = min(ans, max(pre, suf[i/2+1]))
		}
		Fprintln(out, ans)
	}
}

func main() { CF2026B(os.Stdin, os.Stdout) }
