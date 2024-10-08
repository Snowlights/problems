//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF1922C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, m, x, y int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}

		l := make([]int, n)
		r := make([]int, n)
		r[1] = 1
		for i := 1; i < n-1; i++ {
			if a[i]-a[i-1] < a[i+1]-a[i] {
				l[i] = l[i-1] + 1
				r[i+1] = r[i] + a[i+1] - a[i]
			} else {
				l[i] = l[i-1] + a[i] - a[i-1]
				r[i+1] = r[i] + 1
			}
		}
		l[n-1] = l[n-2] + 1

		for Fscan(in, &m); m > 0; m-- {
			Fscan(in, &x, &y)
			if x < y {
				Fprintln(out, r[y-1]-r[x-1])
			} else {
				Fprintln(out, l[x-1]-l[y-1])
			}
		}
	}

}

func main() { CF1922C(os.Stdin, os.Stdout) }
