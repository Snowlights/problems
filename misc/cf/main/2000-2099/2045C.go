//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF2045C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var s, t string
	Fscan(in, &s, &t)
	pos := ['z' + 1]int{}
	for i, c := range t[:len(t)-1] {
		pos[c] = i + 1
	}
	minI := -1
	minL := int(1e9)
	for i := 1; i < len(s); i++ {
		c := s[i]
		if pos[c] > 0 {
			l := i + 1 + len(t) - pos[c]
			if l < minL {
				minL = l
				minI = i
			}
		}
	}
	if minI >= 0 {
		Fprint(out, s[:minI+1], t[pos[s[minI]]:])
	} else {
		Fprint(out, -1)
	}

}

func main() { CF2045C(os.Stdin, os.Stdout) }
