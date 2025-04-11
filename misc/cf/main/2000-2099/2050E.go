//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF2050E(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T int
	var a, b, c []byte
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &a, &b, &c)
		m := len(b)
		f := make([]int, m+1)
		for j, y := range b {
			f[j+1] = f[j] + b2i50(y != c[j])
		}
		for i, x := range a {
			f[0] += b2i50(x != c[i])
			for j, y := range b {
				f[j+1] = min(f[j+1]+b2i50(x != c[i+j+1]), f[j]+b2i50(y != c[i+j+1]))
			}
		}
		Fprintln(out, f[m])
	}

}

func b2i50(b bool) int {
	if b {
		return 1
	}
	return 0
}

func main() { CF2050E(os.Stdin, os.Stdout) }
