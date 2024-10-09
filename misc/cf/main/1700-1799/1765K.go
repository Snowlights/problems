//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF1765K(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, v, s int
	Fscan(in, &n)
	mn := int(1e9)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			Fscan(in, &v)
			s += v
			if j == n-1-i {
				mn = min(mn, v)
			}
		}
	}
	Fprint(out, s-mn)

}

func main() { CF1765K(os.Stdin, os.Stdout) }
