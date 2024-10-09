//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF1759E(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, h int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &h)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		slices.Sort(a)
		f := func(m []int) int {
			sum := h
			for i, v := range a {
				for sum <= v {
					if len(m) == 0 {
						return i
					}
					sum *= m[0]
					m = m[1:]
				}
				sum += v / 2
			}
			return n
		}
		Fprintln(out, max(f([]int{2, 2, 3}), f([]int{2, 3, 2}), f([]int{3, 2, 2})))
	}

}

func main() { CF1759E(os.Stdin, os.Stdout) }
