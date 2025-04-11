//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF327A(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, v, f0 int
	f1, f2 := int(-1e9), int(-1e9)
	Fscan(in, &n)
	for range n {
		Fscan(in, &v)
		f2 = max(f2, f1) + v
		f1 = max(f1, f0) + 1 - v
		f0 += v
	}
	Fprint(out, max(f1, f2))

}

func main() { CF327A(os.Stdin, os.Stdout) }
