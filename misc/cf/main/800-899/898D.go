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

func CF898D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, m, k, cnt, l, ans int
	Fscan(in, &n, &m, &k)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	slices.Sort(a)
	for i, x := range a {
		cnt++
		for x-a[l] >= m {
			if a[l] > 0 {
				cnt--
			}
			l++
		}
		if cnt == k {
			a[i] = 0
			cnt--
			ans++
		}
	}
	Fprint(out, ans)
}

func main() { CF898D(os.Stdin, os.Stdout) }
