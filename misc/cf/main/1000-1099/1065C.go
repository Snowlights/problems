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

func CF1065C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, k, s, ans int
	Fscan(in, &n, &k)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	sort.Sort(sort.Reverse(sort.IntSlice(a)))
	for i := 1; i < n; i++ {
		d := a[i-1] - a[i]
		area := d * i
		if s+area < k {
			s += area
			continue
		}
		d -= (k - s) / i
		ans += 1 + d/(k/i)
		s = d % (k / i) * i
	}
	if s > 0 {
		ans++
	}
	Fprint(out, ans)

}

func main() { CF1065C(os.Stdin, os.Stdout) }
