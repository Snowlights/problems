//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF1031B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n int
	Fscan(in, &n)
	a := make([]int, n-1)
	for i := range a {
		Fscan(in, &a[i])
	}
	b := make([]int, n-1)
	for i := range b {
		Fscan(in, &b[i])
	}

	const mx = 4
	ans := make([]any, n)
	vis := make([][mx]bool, n)
	var dfs func(int, int) bool
	dfs = func(i, j int) bool {
		if i < 0 {
			return true
		}
		if vis[i][j] {
			return false
		}
		vis[i][j] = true
		for k := range mx {
			if k|j == a[i] && k&j == b[i] && dfs(i-1, k) {
				ans[i] = k
				return true
			}
		}
		return false
	}
	for j := range mx {
		ans[n-1] = j
		if dfs(n-2, j) {
			Fprintln(out, "YES")
			Fprintln(out, ans...)
			return
		}
	}
	Fprint(out, "NO")

}

func main() { CF1031B(os.Stdin, os.Stdout) }
