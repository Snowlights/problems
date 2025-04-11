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

func CF1798D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		zero := true
		a := make([]int, n)
		res := make([]int, 0, n)
		for i := range a {
			Fscan(in, &a[i])
			if a[i] != 0 {
				zero = false
			}
		}
		if zero {
			Fprintln(out, "No")
			continue
		}
		sort.Ints(a)
		l, r, sum := 0, n-1, 0
		for l <= r {
			if sum <= 0 {
				sum += a[r]
				res = append(res, a[r])
				r--
			} else {
				sum += a[l]
				res = append(res, a[l])
				l++
			}
		}
		Fprintln(out, "Yes")
		for _, v := range res {
			Fprint(out, v, " ")
		}
		Fprintln(out)
	}

}

func main() { CF1798D(os.Stdin, os.Stdout) }
