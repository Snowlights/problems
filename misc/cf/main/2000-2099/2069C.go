//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF2069C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	const mod = 998244353
	var T, n, v int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		var f1, f2, f3 int
		for range n {
			Fscan(in, &v)
			if v == 1 {
				f1++
			} else if v == 2 {
				f2 = (f2*2 + f1) % mod
			} else {
				f3 += f2
			}
		}
		Fprintln(out, f3%mod)
	}
	
}

func main() { CF2069C(os.Stdin, os.Stdout) }
