//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF1792C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, v int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n+1)
		for i := 0; i < n; i++ {
			Fscan(in, &v)
			a[v] = i
		}
		x := n / 2
		for ; x > 0 && a[x] < a[x+1] && a[n-x] < a[n-x+1]; x-- {
		}
		Fprintln(out, x)
	}

}

func main() { CF1792C(os.Stdin, os.Stdout) }
