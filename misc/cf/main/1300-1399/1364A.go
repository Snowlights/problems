//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF1364A(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, x, v int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &x)
		s := 0
		l, r := -1, -1
		for i := range n {
			Fscan(in, &v)
			s += v
			if v%x != 0 {
				if l < 0 {
					l = i
				}
				r = i
			}
		}
		if s%x != 0 {
			Fprintln(out, n)
		} else if l < 0 {
			Fprintln(out, -1)
		} else {
			Fprintln(out, max(n-l-1, r))
		}
	}

}

func main() { CF1364A(os.Stdin, os.Stdout) }
