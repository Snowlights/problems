//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF2040B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n uint
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		Fprintln(out, 1+bits.Len((n+1)/3))
	}

}

func main() { CF2040B(os.Stdin, os.Stdout) }
