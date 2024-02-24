//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF505B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, m, v, w, c int
	Fscan(in, &n, &m)
	fa := make([][]int, m+1)
	for i := range fa {
		fa[i] = make([]int, n+1)
		for j := range fa[i] {
			fa[i][j] = j
		}
	}

	var find func(a []int, from int) int
	find = func(a []int, from int) int {
		if a[from] == from {
			return from
		}
		a[from] = find(a, a[from])
		return a[from]
	}

	merge := func(a []int, from, to int) {
		from, to = find(a, from), find(a, to)
		if from == to {
			return
		}
		a[from] = to
	}

	for i := 0; i < m; i++ {
		Fscan(in, &v, &w, &c)
		merge(fa[c], v, w)
	}

	Fscan(in, &m)
	for i := 0; i < m; i++ {
		Fscan(in, &v, &w)
		ans := 0
		for _, f := range fa {
			if find(f, v) == find(f, w) {
				ans++
			}
		}
		Fprintln(out, ans)
	}

}

func main() { CF505B(os.Stdin, os.Stdout) }
