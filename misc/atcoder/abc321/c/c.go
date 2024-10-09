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

	var k, s int
	Fscan(in, &k)
	var dfs func(x int)
	ans := make([]int, 0, 1e5)
	dfs = func(x int) {
		s = s*10 + x
		ans = append(ans, s)
		for i := x - 1; i >= 0; i-- {
			dfs(i)
		}
		s /= 10
	}
	for i := 9; i >= 0; i-- {
		dfs(i)
	}
	sort.Ints(ans)
	Fprintln(out, ans[k])
}

func main() { run(os.Stdin, os.Stdout) }
