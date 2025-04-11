//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF1991B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		b := make([]int, n-1)
		ok := true
		for i := range b {
			Fscan(in, &b[i])
			if i > 1 && b[i]&b[i-2]&^b[i-1] > 0 {
				ok = false
			}
		}
		if !ok {
			Fprintln(out, -1)
			continue
		}
		Fprint(out, b[0], " ")
		for i := range n - 2 {
			Fprint(out, b[i]|b[i+1], " ")
		}
		Fprintln(out, b[n-2])
	}

}

func main() { CF1991B(os.Stdin, os.Stdout) }
