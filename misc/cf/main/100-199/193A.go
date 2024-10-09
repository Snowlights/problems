//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF193A(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, m int
	Fscan(in, &n, &m)
	vis, grand, total := make([][]bool, n), make([][]byte, n), 0
	for i := range vis {
		vis[i] = make([]bool, m)
		Fscan(in, &grand[i])
		for _, v := range grand[i] {
			if v == '#' {
				total++
			}
		}
	}
	if total < 3 {
		Fprintln(out, -1)
		return
	}
	resetVis := func() {
		for i := range vis {
			vis[i] = make([]bool, m)
		}
	}
	dir := [][]int{{0, -1}, {0, 1}, {1, 0}, {-1, 0}}
	dfs := func() int {
		cnt, q := 0, [][]int{}
		for i, v := range grand {
			for j, vv := range v {
				if vv == '.' || vis[i][j] {
					continue
				}
				cnt++
				q = [][]int{{i, j}}
				vis[i][j] = true
				for len(q) > 0 {
					tmp := q[0]
					q = q[1:]
					for _, d := range dir {
						x, y := tmp[0]+d[0], tmp[1]+d[1]
						if 0 <= x && x < n && 0 <= y && y < m && !vis[x][y] && grand[x][y] == '#' {
							q = append(q, []int{x, y})
							vis[x][y] = true
						}
					}
				}
			}
		}
		return cnt
	}

	for _, v := range grand {
		for j, vv := range v {
			if vv == '.' {
				continue
			}
			v[j] = '.'
			resetVis()
			if dfs() >= 2 {
				Fprintln(out, 1)
				return
			}
			v[j] = '#'
		}
	}
	Fprintln(out, 2)
	return
}

func main() { CF193A(os.Stdin, os.Stdout) }
