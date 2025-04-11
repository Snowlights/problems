//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF1148B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, m, ta, tb, k int
	Fscan(in, &n, &m, &ta, &tb, &k)
	if k >= min(n, m) {
		Fprint(out, -1)
		return
	}
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	b := make([]int, m)
	for i := range b {
		Fscan(in, &b[i])
	}

	ans, j := 0, 0
	for i, t := range a[:k+1] {
		for j < m && b[j] < t+ta {
			j++
		}
		if j+k-i >= m {
			Fprint(out, -1)
			return
		}
		ans = max(ans, j+k-i)
	}
	Fprint(out, b[ans]+tb)

}

func main() { CF1148B(os.Stdin, os.Stdout) }
