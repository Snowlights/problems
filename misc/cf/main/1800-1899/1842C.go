//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF1842C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, v int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		mx := make([]int, n+1)
		for i := range mx {
			mx[i] = -1e9
		}
		f := 0
		for i := range n {
			Fscan(in, &v)
			f, mx[v] = max(f, mx[v]+i+1), max(mx[v], f-i)
		}
		Fprintln(out, f)
	}

}

func main() { CF1842C(os.Stdin, os.Stdout) }
