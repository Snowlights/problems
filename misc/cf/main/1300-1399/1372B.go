//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF1372B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n int
o:
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		for d := 2; d*d <= n; d++ {
			if n%d == 0 {
				Fprintln(out, n/d, n-n/d)
				continue o
			}
		}
		Fprintln(out, 1, n-1)
	}

}

func main() { CF1372B(os.Stdin, os.Stdout) }
