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

	var n, m, from, to, ans int
	Fscan(in, &n, &m)
	fa := make([]int, n+1)
	for i := range fa {
		fa[i] = i
	}
	var find func(int) int
	find = func(x int) int {
		if fa[x] != x {
			fa[x] = find(fa[x])
		}
		return fa[x]
	}
	merge := func(from, to int) {
		from, to = find(from), find(to)
		if from == to {
			return
		}
		fa[from] = to
	}

	for i := 0; i < m; i++ {
		Fscan(in, &from, &to)
		if find(from) == find(to) {
			ans++
		} else {
			merge(from, to)
		}
	}

	Fprintln(out, ans)

}

func main() { run(os.Stdin, os.Stdout) }
