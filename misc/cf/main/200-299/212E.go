//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF212E(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n int
	Fscan(in, &n)
	g := make([][]int, n)
	for range n - 1 {
		var v, w int
		Fscan(in, &v, &w)
		v--
		w--
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}

	has := make([]bool, n-1)
	var dfs func(int, int) int
	dfs = func(v, fa int) int {
		size := 1
		f := make([]bool, n-1)
		f[0] = true
		for _, w := range g[v] {
			if w == fa {
				continue
			}
			sz := dfs(w, v)
			for j := n - 2; j >= sz; j-- {
				f[j] = f[j] || f[j-sz]
			}
			size += sz
		}
		sz := n - size
		for j := n - 2; j >= sz; j-- {
			f[j] = f[j] || f[j-sz]
		}
		for i, b := range f {
			if b {
				has[i] = true
			}
		}
		return size
	}
	dfs(0, -1)

	has[0] = false
	tot := 0
	for _, v := range has {
		if v {
			tot++
		}
	}
	Fprintln(out, tot)
	for i, b := range has {
		if b {
			Fprintln(out, i, n-1-i)
		}
	}

}

func main() { CF212E(os.Stdin, os.Stdout) }
