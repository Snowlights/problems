//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF903F(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n int
	Fscan(in, &n)
	a := [4]int32{}
	s := [4]string{}
	Fscan(in, &a[0], &a[1], &a[2], &a[3], &s[0], &s[1], &s[2], &s[3])

	memo := make([][4][1 << 4][1 << 4][1 << 4]int32, n)
	var dfs func(int, int, int, int, int) int32
	dfs = func(j, i, cur, pre, pre2 int) int32 {
		if j < 0 {
			if pre > 0 || pre2 > 0 {
				return a[3]
			}
			return 0
		}
		if i > 3 {
			if pre2 > 0 {
				return dfs(j-2, 0, 0, 0, 0) + a[3]
			}
			return dfs(j-1, 0, 0, cur, pre)
		}
		p := &memo[j][i][cur][pre][pre2]
		if *p > 0 {
			return *p - 1
		}
		res := min(
			dfs(j, i+1, cur<<1|int(s[i][j]>>2&1^1), pre, pre2),
			dfs(j, i+1, cur<<1, pre, pre2)+a[0],
			dfs(j, i+2, cur<<min(4-i, 2), pre&^(3<<max(2-i, 0)), pre2)+a[1],
			dfs(j, i+3, cur<<min(4-i, 3), pre&^(7<<max(1-i, 0)), pre2&^(7<<max(1-i, 0)))+a[2],
		)
		*p = res + 1
		return res
	}
	Fprint(out, dfs(n-1, 0, 0, 0, 0))
}

func main() { CF903F(os.Stdin, os.Stdout) }
