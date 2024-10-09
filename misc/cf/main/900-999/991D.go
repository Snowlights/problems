//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF991D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var s1, s2 []byte
	Fscan(in, &s1, &s2)

	shape := [][3]int{{0, 0, 0}, {3, 1, 1}, {3, 2, 1}, {1, 3, 1}, {2, 3, 1}}
	f := [4]int{-1e9, -1e9, -1e9, 0}
	for i, b := range s1 {
		x := int(b>>6 | s2[i]>>6<<1)
		nf := [4]int{-1e9, -1e9, -1e9, -1e9}
		for _, p := range shape {
			if p[1]&x > 0 {
				continue
			}
			for pre := 0; pre < 4; pre++ {
				if p[0]&pre == 0 {
					nf[p[1]|x] = max(nf[p[1]|x], f[pre]+p[2])
				}
			}
		}
		f = nf
	}
	Fprint(out, max(max(max(f[0], f[1]), f[2]), f[3]))
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func CF991DV2(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var s1, s2 []byte
	Fscan(in, &s1, &s2)
	n, r, p := len(s1), 0, 0
	for i := 0; i < n; i++ {
		c := 2
		if s1[i] != s2[i] {
			c = 1
		} else {
			if s1[i] == 'X' {
				c = 0
			}
		}
		p += c
		if p < 3 {
			p = c
		} else {
			r++
			p -= 3
		}
	}
	Fprintln(out, r)
}

func main() { CF991DV2(os.Stdin, os.Stdout) }
