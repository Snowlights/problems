//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF1717C(_r io.Reader, _w io.Writer) {
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
		b := make([]int, n)
		for i := range b {
			Fscan(in, &b[i])
		}
		for i, v := range b {
			if a[i] > v || a[i] < v && v-b[(i+1)%n] > 1 {
				Fprintln(out, "NO")
				continue o
			}
		}
		Fprintln(out, "YES")
	}

}

func main() { CF1717C(os.Stdin, os.Stdout) }
