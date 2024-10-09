//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"strings"
)

func CF746D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, a, b, k int
	Fscan(in, &n, &k, &a, &b)
	cg, cb := "G", "B"
	if a < b {
		a, b, cg, cb = b, a, cb, cg
	}
	b++
	if (a-1)/b >= k {
		Fprint(out, "NO")
		return
	}
	base := cb + strings.Repeat(cg, a/b)
	Fprint(out, strings.Repeat(base, b-a%b)[1:], strings.Repeat(base+cg, a%b))

}

func main() { CF746D(os.Stdin, os.Stdout) }
