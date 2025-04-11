//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"slices"
)

func CF709B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}
	var n, x int
	Fscan(in, &n, &x)
	if n == 1 {
		Fprint(out, 0)
		return
	}
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	slices.Sort(a)
	f := func(a []int) int {
		return a[n-2] - a[0] + min(abs(x-a[0]), abs(x-a[n-2]))
	}
	Fprint(out, min(f(a[:n-1]), f(a[1:])))

}

func main() { CF709B(os.Stdin, os.Stdout) }
