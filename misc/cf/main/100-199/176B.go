//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF176B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	const mod = 1_000_000_007
	var s, t string
	var k, c, f int64
	Fscan(in, &s, &t, &k)
	n := int64(len(s))
	ss := s + s
	for i := int64(0); i < n; i++ {
		if ss[i:i+n] == t {
			c++
		}
	}
	if s == t {
		f = 1
	}
	g := f ^ 1
	for ; k > 0; k-- {
		f, g = (f*(c-1)+g*c)%mod,
			(f*(n-c)+g*(n-c-1))%mod
	}
	Fprint(out, f)

}

func main() { CF176B(os.Stdin, os.Stdout) }
