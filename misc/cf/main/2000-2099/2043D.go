//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF2043D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	gcd := func(a, b int) int {
		for a != 0 {
			a, b = b%a, a
		}
		return b
	}
	var T, l, r, g int
o:
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &l, &r, &g)
		l = (l-1)/g + 1
		r /= g
		for d := range r - l + 1 {
			for i := range d + 1 {
				if gcd(l+i, r-(d-i)) == 1 {
					Fprintln(out, (l+i)*g, (r-(d-i))*g)
					continue o
				}
			}
		}
		Fprintln(out, -1, -1)
	}

}

func main() { CF2043D(os.Stdin, os.Stdout) }
