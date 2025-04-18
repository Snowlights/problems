//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF2046B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, v int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		var a, b []int
		minB := math.MaxInt
		for range n {
			Fscan(in, &v)
			for len(a) > 0 && v < a[len(a)-1] {
				t := a[len(a)-1]
				b = append(b, t+1)
				minB = min(minB, t+1)
				a = a[:len(a)-1]
			}
			if v > minB {
				b = append(b, v+1)
			} else {
				a = append(a, v)
			}
		}
		slices.Sort(b)
		for _, v := range append(a, b...) {
			Fprint(out, v, " ")
		}
		Fprintln(out)
	}
}

func main() { CF2046B(os.Stdin, os.Stdout) }
