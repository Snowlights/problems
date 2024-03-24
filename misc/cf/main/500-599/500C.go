//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF500C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, m, ans int
	Fscan(in, &n, &m)
	w := make([]int, n+1)
	for i := 1; i <= n; i++ {
		Fscan(in, &w[i])
	}
	b := make([]int, m+1)
	last := make([]int, n+1)
	vis := make([]int, n+1)
	for i := 1; i <= m; i++ {
		Fscan(in, &b[i])
		for _, v := range b[last[b[i]]+1 : i] {
			if vis[v] != i {
				vis[v] = i
				ans += w[v]
			}
		}
		last[b[i]] = i
	}
	Fprint(out, ans)
	
}

func main() { CF500C(os.Stdin, os.Stdout) }
