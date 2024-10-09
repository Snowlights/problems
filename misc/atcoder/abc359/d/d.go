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

	const mod = 998244353
	var n, k int
	var s string
	Fscan(in, &n, &k, &s)

	pal := make([]bool, 1<<k)
o:
	for i := range pal {
		for j := 0; j < k/2; j++ {
			if i>>j&1 != i>>(k-1-j)&1 {
				continue o
			}
		}
		pal[i] = true
	}

	mask := 1<<(k-1) - 1
	memo := make([][]int, n)
	for i := range memo {
		memo[i] = make([]int, 1<<k)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}
	var dfs func(int, int) int
	dfs = func(i, j int) (res int) {
		if i < 0 {
			return 1
		}
		p := &memo[i][j]
		if *p != -1 {
			return *p
		}
		for b := 0; b < 2; b++ {
			if s[i] != '?' && int(s[i]&1) != b {
				continue
			}
			t := j<<1 | b
			if i > n-k || !pal[t] {
				res += dfs(i-1, t&mask)
			}
		}
		res %= mod
		*p = res
		return
	}
	Fprint(out, dfs(n-1, 0))

}

func main() { run(os.Stdin, os.Stdout) }
