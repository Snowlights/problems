//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF1893B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, m int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		b := make([]int, m)
		for i := range b {
			Fscan(in, &b[i])
		}
		slices.SortFunc(b, func(a, b int) int { return b - a })

		j := 0
		for _, v := range a {
			for ; j < m && b[j] >= v; j++ {
				Fprint(out, b[j], " ")
			}
			Fprint(out, v, " ")
		}
		for _, v := range b[j:] {
			Fprint(out, v, " ")
		}
		Fprintln(out)
	}

}

func main() { CF1893B(os.Stdin, os.Stdout) }
