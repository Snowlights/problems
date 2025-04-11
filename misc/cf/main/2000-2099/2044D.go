//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF2044D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, v int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		vis := make([]bool, n+1)
		for range n {
			Fscan(in, &v)
			if !vis[v] {
				vis[v] = true
				Fprint(out, v, " ")
			}
		}
		for i := 1; i <= n; i++ {
			if !vis[i] {
				Fprint(out, i, " ")
			}
		}
		Fprintln(out)
	}
}

func main() { CF2044D(os.Stdin, os.Stdout) }
