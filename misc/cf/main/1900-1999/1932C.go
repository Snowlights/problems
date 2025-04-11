//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF1932C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, mod int
	var s string
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &mod)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		Fscan(in, &s)

		ans := make([]any, n)
		mul := 1
		r := strings.Count(s, "L")
		l := r - 1
		for i := n - 1; i >= 0; i-- {
			if s[i] == 'L' {
				mul = mul * a[l] % mod
				l--
			} else {
				mul = mul * a[r] % mod
				r++
			}
			ans[i] = mul
		}
		Fprintln(out, ans...)
	}

}

func main() { CF1932C(os.Stdin, os.Stdout) }
