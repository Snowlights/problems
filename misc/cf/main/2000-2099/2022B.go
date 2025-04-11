//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF2022B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, x, v int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &x)
		s, mx := 0, 0
		for range n {
			Fscan(in, &v)
			s += v
			mx = max(mx, v)
		}
		Fprintln(out, max((s-1)/x+1, mx))
	}

}

func main() { CF2022B(os.Stdin, os.Stdout) }
