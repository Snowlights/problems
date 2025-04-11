//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF262B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}
	var n, k, s int
	Fscan(in, &n, &k)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}

	slices.Sort(a)
	mn := int(1e9)
	for i, v := range a {
		mn = min(mn, abs(v))
		if k > 0 && v < 0 {
			a[i] = -v
			k--
		}
		s += a[i]
	}

	if k%2 > 0 {
		s -= mn * 2
	}
	Fprint(out, s)

}

func main() { CF262B(os.Stdin, os.Stdout) }
