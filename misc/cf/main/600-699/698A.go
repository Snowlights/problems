//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"math"
	"os"
)

func CF698A(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	const inf = math.MaxInt32

	var n, v, f0, f1, f2 int
	Fscan(in, &n)
	for i := 0; i < n; i++ {
		Fscan(in, &v)
		g1, g2 := inf, inf
		if v&1 > 0 {
			g1 = min(f0, f2)
		}
		if v>>1 > 0 {
			g2 = min(f0, f1)
		}
		f0 = min(min(f0, f1), f2) + 1
		f1, f2 = g1, g2
	}
	Fprintln(out, min(min(f0, f1), f2))
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() { CF698A(os.Stdin, os.Stdout) }
