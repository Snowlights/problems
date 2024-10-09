//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF812C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, s, mn int
	Fscan(in, &n, &s)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	ans := sort.Search(n, func(k int) bool {
		k++
		b := slices.Clone(a)
		for i := range b {
			b[i] += (i + 1) * k
		}
		slices.Sort(b)
		tot := 0
		for _, v := range b[:k] {
			tot += v
		}
		if tot > s {
			return true
		}
		mn = tot
		return false
	})
	Fprintln(out, ans, mn)
	
}

func main() { CF812C(os.Stdin, os.Stdout) }
