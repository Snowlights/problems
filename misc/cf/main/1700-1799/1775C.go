//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF1775C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, x int
o:
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &x)
		if n == x {
			Fprintln(out, n)
			continue
		}
		for n > x {
			lb := n & -n
			n ^= lb
			if n == x && n&(lb<<1) == 0 {
				Fprintln(out, n|lb<<1)
				continue o
			}
		}
		Fprintln(out, -1)
	}
	
}

func main() { CF1775C(os.Stdin, os.Stdout) }
