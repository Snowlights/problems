//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF1713B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n int
o:
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		a = slices.Compact(a)
		for i := 1; i < len(a)-1; i++ {
			if a[i-1] > a[i] && a[i] < a[i+1] {
				Fprintln(out, "NO")
				continue o
			}
		}
		Fprintln(out, "YES")
	}
	
}

func main() { CF1713B(os.Stdin, os.Stdout) }
