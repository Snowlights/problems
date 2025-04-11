//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF1870B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, m, v int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		var or, x1, x2 int
		for range m {
			Fscan(in, &v)
			or |= v
		}
		for _, v := range a {
			x1 ^= v
			x2 ^= v | or
		}
		Fprintln(out, min(x1, x2), max(x1, x2))
	}

}

func main() { CF1870B(os.Stdin, os.Stdout) }
