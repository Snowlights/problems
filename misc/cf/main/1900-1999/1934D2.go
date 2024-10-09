//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
	"os"
)

func CF1934D2(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, p, q uint
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		if bits.OnesCount(n)%2 > 0 {
			Fprintln(out, "second")
			Fscan(in, &p, &q)
			if bits.OnesCount(q)%2 > 0 {
				n = p
			} else {
				n = q
			}
		} else {
			Fprintln(out, "first")
		}
		for n > 0 {
			h := uint(1) << (bits.Len(n) - 1)
			Fprintln(out, n^h, h)
			Fscan(in, &p, &q)
			if bits.OnesCount(q)%2 > 0 {
				n = p
			} else {
				n = q
			}
		}
	}

}

func main() { CF1934D2(os.Stdin, os.Stdout) }
