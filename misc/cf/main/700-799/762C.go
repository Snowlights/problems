//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF762C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var s, t string
	Fscan(in, &s, &t)
	n, m := len(s), len(t)
	suf := make([]int, n+1)
	suf[n] = m
	for i, j := n-1, m-1; i >= 0; i-- {
		if s[i] == t[j] {
			j--
		}
		if j < 0 {
			Fprint(out, t)
			return
		}
		suf[i] = j + 1
	}

	l, r := 0, suf[0]
	j := 0
	for i := range s {
		if s[i] == t[j] {
			j++
		}
		if suf[i+1]-j < r-l {
			l, r = j, suf[i+1]
		}
	}
	if r-l == m {
		Fprint(out, "-")
	} else {
		Fprint(out, t[:l], t[r:])
	}

}

func main() { CF762C(os.Stdin, os.Stdout) }
