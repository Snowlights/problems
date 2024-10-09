package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, m int
	Fscan(in, &n, &m)
	grid := make([]string, n)
	for i := range grid {
		Fscan(in, &grid[i])
	}

	s := "snuke"
	dir := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	vis, dp := make([][]bool, n), make([][]int, n)
	for i := range vis {
		vis[i] = make([]bool, m)
		dp[i] = make([]int, m)
	}

	var dfs func(i, j, idx int) bool
	dfs = func(i, j, idx int) bool {
		if i == n-1 && j == m-1 {
			return true
		}
		if dp[i][j] == 1 {
			return false
		}
		for _, d := range dir {
			if x, y := d[0]+i, d[1]+j; 0 <= x && x < n && 0 <= y && y < m && !vis[x][y] && grid[x][y] == s[idx] {
				vis[x][y] = true
				if dfs(x, y, (idx+1)%5) {
					return true
				}
				vis[x][y] = false
			}
		}
		dp[i][j] = 1
		return false
	}

	if grid[0][0] != s[0] || !dfs(0, 0, 1) {
		Fprintln(out, "No")
	} else {
		Fprintln(out, "Yes")
	}

}

func main() { run(os.Stdin, os.Stdout) }
