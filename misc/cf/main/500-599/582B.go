//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF582B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, t, d int
	Fscan(in, &n, &t)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	f := []int{}
	for i := 0; i < n && t > 0; i++ {
		m := len(f)
		for _, x := range a {
			j := sort.SearchInts(f, x+1)
			if j < len(f) {
				f[j] = x
			} else {
				f = append(f, x)
			}
		}
		d = len(f) - m
		t--
	}
	Fprint(out, len(f)+t*d)

}

func main() { CF582B(os.Stdin, os.Stdout) }
