//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF2031D(_r io.Reader, _w io.Writer) {
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

		suf := make([]int, n)
		suf[n-1] = a[n-1]
		for i := n - 2; i >= 0; i-- {
			suf[i] = min(suf[i+1], a[i])
		}

		for i := 0; i < n; {
			st := i
			mx := a[i]
			for i++; i < n && mx > suf[i]; i++ {
				mx = max(mx, a[i])
			}
			for range i - st {
				Fprint(out, mx, " ")
			}
		}
		Fprintln(out)
	}

}

func main() { CF2031D(os.Stdin, os.Stdout) }
