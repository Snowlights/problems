//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF371D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, m, op, b, w int
	Fscan(in, &n)
	required, fa := make([]int, n), make([]int, n)
	water := make([]int, n)
	for i := range required {
		Fscan(in, &required[i])
		fa[i] = i
	}
	var find func(int) int
	find = func(x int) int {
		if x >= n {
			return x
		}
		if fa[x] != x {
			fa[x] = find(fa[x])
		}
		return fa[x]
	}
	Fscan(in, &m)
	for i := 0; i < m; i++ {
		Fscan(in, &op)
		switch op {
		case 1:
			Fscan(in, &b, &w)
			b--
			for j := find(b); j < n; j = find(j) {
				if required[j]-water[j] >= w {
					water[j] += w
					break
				}
				w -= required[j] - water[j]
				water[j] = required[j]
				fa[j] = j + 1
			}
		case 2:
			Fscan(in, &b)
			b--
			Fprintln(out, water[b])
		}
	}

}

func main() { CF371D(os.Stdin, os.Stdout) }
