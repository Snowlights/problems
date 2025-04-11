//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF2027C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, v int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &v)
		g := map[int][]int{}
		for i := 1; i < n; i++ {
			Fscan(in, &v)
			v += i
			g[v] = append(g[v], v+i)
		}

		ans := n
		vis := map[int]bool{}
		var dfs func(int)
		dfs = func(v int) {
			vis[v] = true
			ans = max(ans, v)
			for _, w := range g[v] {
				if !vis[w] {
					dfs(w)
				}
			}
		}
		dfs(n)
		Fprintln(out, ans)
	}

}

func main() { CF2027C(os.Stdin, os.Stdout) }
