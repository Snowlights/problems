//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF1906B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n int
	var a string
	f := func() (c1 int) {
		Fscan(in, &a)
		s := 0
		for _, b := range a {
			s ^= int(b & 1)
			c1 += s
		}
		return
	}
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		c1 := f()
		c2 := f()
		if c1 == c2 || c1 == n+1-c2 {
			Fprintln(out, "YES")
		} else {
			Fprintln(out, "NO")
		}
	}

}

func main() { CF1906B(os.Stdin, os.Stdout) }
