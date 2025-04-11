//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF1632B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		hb := 1 << (bits.Len(uint(n-1)) - 1)
		for i := 1; i < hb; i++ {
			Fprint(out, i, " ")
		}
		Fprint(out, "0 ", hb)
		for i := hb + 1; i < n; i++ {
			Fprint(out, " ", i)
		}
		Fprintln(out)
	}
}

func main() { CF1632B(os.Stdin, os.Stdout) }
