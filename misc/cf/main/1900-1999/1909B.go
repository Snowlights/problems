//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF1909B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, v int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		or, and := 0, -1
		for range n {
			Fscan(in, &v)
			or |= v
			and &= v
		}
		or ^= and
		Fprintln(out, or&-or<<1)
	}

}

func main() { CF1909B(os.Stdin, os.Stdout) }
