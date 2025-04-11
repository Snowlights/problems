//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF1974E(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, x, c, h int
	f := [5e4 + 1]int{}
	for i := 1; i < len(f); i++ {
		f[i] = 1e18
	}
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &x)
		s := 0
		for i := 0; i < n; i++ {
			Fscan(in, &c, &h)
			s += h
			for j := s; j >= h; j-- {
				if f[j-h] <= i*x-c {
					f[j] = min(f[j], f[j-h]+c)
				}
			}
		}
		for i := s; i >= 0; i-- {
			if f[i] < 1e18 {
				Fprintln(out, i)
				break
			}
		}
		for i := s; i > 0; i-- {
			f[i] = 1e18
		}
		f[0] = 0
	}

}

func main() { CF1974E(os.Stdin, os.Stdout) }
