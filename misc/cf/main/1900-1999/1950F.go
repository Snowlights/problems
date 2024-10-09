//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF1950F(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, a, b, c int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &a, &b, &c)
		if a+1 != c {
			Fprintln(out, -1)
			continue
		}
		h := bits.Len(uint(a))
		b -= 1<<h - 1 - a
		if b > 0 {
			h += (b + a) / (a + 1)
		}
		Fprintln(out, h)
	}

}

func main() { CF1950F(os.Stdin, os.Stdout) }
