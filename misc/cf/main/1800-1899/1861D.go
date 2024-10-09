//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF1861D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a, ans := make([]int, n), 0
		for i := range a {
			Fscan(in, &a[i])
		}
		suf := make([]int, n)
		for i := n - 2; i >= 0; i-- {
			suf[i] = suf[i+1]
			if a[i] >= a[i+1] {
				suf[i]++
			}
		}
		ans = suf[0]
		pre := 1
		for i := 1; i < n; i++ {
			ans = min(ans, pre+suf[i])
			if a[i] >= a[i-1] {
				pre++
			}
		}
		ans = min(ans, pre)
		Fprintln(out, ans)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() { CF1861D(os.Stdin, os.Stdout) }
