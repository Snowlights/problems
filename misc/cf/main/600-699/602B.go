//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF602B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n int
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}

	cnt, l, ans := make(map[int]int), 0, 0
	for r, v := range a {
		cnt[v]++
		for len(cnt) > 2 {
			cnt[a[l]]--
			if cnt[a[l]] == 0 {
				delete(cnt, a[l])
			}
			l++
		}
		ans = max(ans, r-l+1)
	}

	Fprintln(out, ans)
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func main() { CF602B(os.Stdin, os.Stdout) }
