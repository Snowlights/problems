//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF1900C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n int
	var s string
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &s)
		g := make([][2]int, n)
		for i := range n {
			Fscan(in, &g[i][0], &g[i][1])
		}
		var dfs func(int) int
		dfs = func(v int) int {
			l, r := g[v][0]-1, g[v][1]-1
			if l < 0 && r < 0 {
				return 0
			}
			res := int(1e9)
			if l >= 0 {
				res = dfs(l)
				if s[v] != 'L' {
					res++
				}
			}
			if r >= 0 {
				res2 := dfs(r)
				if s[v] != 'R' {
					res2++
				}
				res = min(res, res2)
			}
			return res
		}
		Fprintln(out, dfs(0))
	}

}

func main() { CF1900C(os.Stdin, os.Stdout) }
