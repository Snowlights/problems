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

func CF2065D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, m, v int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m)
		ans := 0
		a := make([]int, n)
		for i := range a {
			for j := m; j > 0; j-- {
				Fscan(in, &v)
				ans += v * j
				a[i] += v
			}
		}
		slices.Sort(a)
		for i, s := range a {
			ans += s * i * m
		}
		Fprintln(out, ans)
	}

}

func main() { CF2065D(os.Stdin, os.Stdout) }
