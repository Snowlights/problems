//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF961B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, k, maxS1 int
	Fscan(in, &n, &k)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	t := make([]int, n)
	for i := range t {
		Fscan(in, &t[i])
		t[i] ^= 1
	}
	s := [2]int{}
	for i, c := range a {
		s[t[i]] += c
		if i < k-1 {
			continue
		}
		maxS1 = max(maxS1, s[1])
		if t[i-k+1] > 0 {
			s[1] -= a[i-k+1]
		}
	}
	Fprint(out, s[0]+maxS1)

}

func main() { CF961B(os.Stdin, os.Stdout) }
