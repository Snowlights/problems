//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF2033F(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	const mod = 1_000_000_007
	var T, n, m int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m)
		if m == 1 {
			Fprintln(out, n%mod)
			continue
		}
		f0, f1 := 1, 1
		for i := 3; ; i++ {
			nf := (f0 + f1) % m
			if nf == 0 {
				Fprintln(out, n%mod*i%mod)
				break
			}
			f0 = f1
			f1 = nf
		}
	}

}

func main() { CF2033F(os.Stdin, os.Stdout) }
