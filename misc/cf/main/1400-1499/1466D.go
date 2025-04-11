//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF1466D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, v int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n)
		s := 0
		for i := range a {
			Fscan(in, &a[i])
			s += a[i]
		}
		b := make([]int, 0, n-2)
		vis := make([]bool, n)
		for range n*2 - 2 {
			Fscan(in, &v)
			v--
			if !vis[v] {
				vis[v] = true
			} else {
				b = append(b, a[v])
			}
		}
		slices.SortFunc(b, func(a, b int) int { return b - a })
		for _, v := range b {
			Fprint(out, s, " ")
			s += v
		}
		Fprintln(out, s)
	}

}

func main() { CF1466D(os.Stdin, os.Stdout) }
