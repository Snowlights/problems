//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF1840F(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	type tuple struct{ x, y, c int }
	dir3 := []tuple{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}}

	var T, n, m, r int
o:
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m, &r)
		type pair struct{ i, t int }
		die := [2]map[pair]bool{{}, {}}
		for i := 0; i < r; i++ {
			var t, tp, c int
			Fscan(in, &t, &tp, &c)
			die[tp-1][pair{c, t}] = true
		}
		vis := make([][][101]bool, n+1)
		for i := range vis {
			vis[i] = make([][101]bool, m+1)
		}
		vis[0][0][0] = true
		q := []tuple{{}}
		for len(q) > 0 {
			p := q[0]
			q = q[1:]
			for _, d := range dir3 {
				x, y, c := p.x+d.x, p.y+d.y, p.c+d.c
				if x <= n && y <= m && c <= r && !vis[x][y][c] &&
					!die[0][pair{x, x + y + c}] && !die[1][pair{y, x + y + c}] {
					if x == n && y == m {
						Fprintln(out, x+y+c)
						continue o
					}
					vis[x][y][c] = true
					q = append(q, tuple{x, y, c})
				}
			}
		}
		Fprintln(out, -1)
	}

}

func main() { CF1840F(os.Stdin, os.Stdout) }
