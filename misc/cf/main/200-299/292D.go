//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF292D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, m, k, from, to int
	Fscan(in, &n, &m)
	fa := [500]int{}
	var find func(x int) int
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
	for i := range fa {
		fa[i] = i
	}

	pre, suf := make([][500]int, m+1), make([][500]int, m+1)
	edges := make([]struct{ x, y int }, m)
	pre[0] = fa
	for i := 0; i < m; i++ {
		Fscan(in, &edges[i].x, &edges[i].y)
		edges[i].x--
		edges[i].y--
		merge(edges[i].x, edges[i].y)
		pre[i+1] = fa
	}
	for i := range fa {
		fa[i] = i
	}
	suf[m] = fa
	for i := m - 1; i >= 0; i-- {
		merge(edges[i].x, edges[i].y)
		suf[i] = fa
	}

	for Fscan(in, &k); k > 0; k-- {
		Fscan(in, &from, &to)
		fa = pre[from-1]
		for i, f := range suf[to][:n] {
			merge(fa[i], f)
		}
		cnt := 0
		for i, v := range fa[:n] {
			if i == v {
				cnt++
			}
		}
		Fprintln(out, cnt)
	}

}

func main() { CF292D(os.Stdin, os.Stdout) }
