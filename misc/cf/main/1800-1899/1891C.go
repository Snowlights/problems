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

func CF1891C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n)
		s := 0
		for i := range a {
			Fscan(in, &a[i])
			s += a[i]
		}
		slices.Sort(a)

		left := (s + 1) / 2
		ans := left
		for i, v := range a {
			if left < v {
				ans += n - i
				break
			}
			left -= v
		}
		Fprintln(out, ans)
	}

}

func main() { CF1891C(os.Stdin, os.Stdout) }
