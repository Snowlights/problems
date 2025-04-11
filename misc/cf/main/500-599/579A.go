//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF579A(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n uint
	Fscan(in, &n)
	Fprint(out, bits.OnesCount(n))

}

func main() { CF579A(os.Stdin, os.Stdout) }
