//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF1974D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n int
	var s []byte
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &s)
		w := ['X']int{'W': 1, 'E': 1}
		for i, b := range s {
			s[i] = "RH"[w[b]]
			w[b] ^= 1
		}
		c := bytes.Count(s, []byte{'R'})
		if c == 0 || c == n || w['N'] != w['S'] || w['W'] != w['E'] {
			Fprintln(out, "NO")
		} else {
			Fprintf(out, "%s\n", s)
		}
	}

}

func main() { CF1974D(os.Stdin, os.Stdout) }
