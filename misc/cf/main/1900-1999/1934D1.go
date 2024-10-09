//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF1934D1(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, m uint
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m)
		if n^m < n {
			Fprintln(out, 1, n, m)
			continue
		}
		x := uint(1)<<bits.Len(n^1<<(bits.Len(n)-1)) - 1
		if x < m {
			Fprintln(out, -1)
		} else {
			Fprintln(out, 2, n, x, m)
		}
	}

}

func main() { CF1934D1(os.Stdin, os.Stdout) }
