//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF1706C(_r io.Reader, _w io.Writer) {
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
		if n%2 > 0 {
			ans := 0
			for i := 1; i < n; i += 2 {
				ans += max(max(a[i-1], a[i+1])-a[i]+1, 0)
			}
			Fprintln(out, ans)
			continue
		}
		suf := 0
		for i := n - 2; i > 0; i -= 2 {
			suf += max(max(a[i-1], a[i+1])-a[i]+1, 0)
		}
		ans := suf
		pre := 0
		for i := 1; i < n-2; i += 2 {
			suf -= max(max(a[i], a[i+2])-a[i+1]+1, 0)
			pre += max(max(a[i-1], a[i+1])-a[i]+1, 0)
			ans = min(ans, pre+suf)
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

func main() { CF1706C(os.Stdin, os.Stdout) }
