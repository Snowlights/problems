//go:build main
// +build main

package main

import (
	"bufio"
	. "cmp"
	. "fmt"
	"io"
	"os"
)

func CF2029C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, v int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		f := [3]int{0, -1e9, -1e9}
		for range n {
			Fscan(in, &v)
			f[2] = max(f[2]+Compare(v, f[2]), f[1]+Compare(v, f[1]))
			f[1] = max(f[1], f[0])
			f[0] += Compare(v, f[0])
		}
		Fprintln(out, max(f[1], f[2]))
	}

}

func main() { CF2029C(os.Stdin, os.Stdout) }
