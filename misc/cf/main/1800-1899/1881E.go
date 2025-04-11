//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF1881E(_r io.Reader, _w io.Writer) {
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
		f := make([]int, n+2)
		f[n+1] = 1e9
		for i := n - 1; i >= 0; i-- {
			f[i] = min(f[i+1]+1, f[min(i+a[i]+1, n+1)])
		}
		Fprintln(out, f[0])
	}

}

func main() { CF1881E(os.Stdin, os.Stdout) }
