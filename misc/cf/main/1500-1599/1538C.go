//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"sort"
)

func CF1538C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, l, r, ans int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &l, &r)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		sort.Ints(a)
		ans = 0
		for i, v := range a {
			L, R := sort.SearchInts(a[:i], l-v), sort.SearchInts(a[:i], r+1-v)
			ans += R - L
		}
		Fprintln(out, ans)
	}

}

func main() { CF1538C(os.Stdin, os.Stdout) }
