//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF1898B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		m, ans := a[n-1], 0
		for i := n - 2; i >= 0; i-- {
			k := (a[i] - 1) / m
			ans += k
			m = a[i] / (k + 1)
		}
		Fprintln(out, ans)
	}

}

func main() { CF1898B(os.Stdin, os.Stdout) }
