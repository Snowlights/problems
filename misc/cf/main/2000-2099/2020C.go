//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF2020C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, b, c, d int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &b, &c, &d)
		a := b ^ d
		if a|b-a&c == d {
			Fprintln(out, a)
		} else {
			Fprintln(out, -1)
		}
	}
}

func main() { CF2020C(os.Stdin, os.Stdout) }
