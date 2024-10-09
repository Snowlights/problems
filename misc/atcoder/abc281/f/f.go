package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"sort"
)

func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n int
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	sort.Ints(a)
	var dfs func(l, r, t int) int
	dfs = func(l, r, t int) int {
		if l == r || t < 0 {
			return 0
		}
		p := l
		for i := l; i <= r; i++ {
			if (a[i]>>t)&1 == 1 {
				p = i
				break
			}
		}
		if p == l {
			return dfs(l, r, t-1)
		}
		return min(dfs(l, p-1, t-1), dfs(p, r, t-1)) | (1 << t)
	}

	Fprintln(out, dfs(0, n-1, 30))
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() { run(os.Stdin, os.Stdout) }
