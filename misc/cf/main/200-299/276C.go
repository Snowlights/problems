//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF276C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, q, l, r, ans int
	Fscan(in, &n, &q)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	slices.Sort(a)

	diff := make([]int, n+1)
	for range q {
		Fscan(in, &l, &r)
		diff[l-1]++
		diff[r]--
	}
	for i := range n - 1 {
		diff[i+1] += diff[i]
	}
	slices.Sort(diff[:n])

	for i, v := range a {
		ans += v * diff[i]
	}
	Fprint(out, ans)

}

func main() { CF276C(os.Stdin, os.Stdout) }
