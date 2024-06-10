//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF1029C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	type pair struct{ l, r int }
	merge := func(a, b pair) pair { return pair{max(a.l, b.l), min(a.r, b.r)} }

	var n, ans int
	Fscan(in, &n)
	a := make([]pair, n)
	for i := range a {
		Fscan(in, &a[i].l, &a[i].r)
	}

	suf := make([]pair, n+1)
	suf[n].r = 1e9
	for i := n - 1; i > 0; i-- {
		suf[i] = merge(suf[i+1], a[i])
	}

	pre := pair{0, 1e9}
	for i, p := range a {
		m := merge(pre, suf[i+1])
		ans = max(ans, m.r-m.l)
		pre = merge(pre, p)
	}
	Fprintln(out, ans)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() { CF1029C(os.Stdin, os.Stdout) }
