//go:build main
// +build main

package main

import (
	"bufio"
	"bytes"
	. "fmt"
	"io"
	"os"
	"slices"
)

func CF1227G(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n int
	Fscan(in, &n)
	type pair struct{ v, i int }
	a := make([]pair, n)
	for i := range a {
		Fscan(in, &a[i].v)
		a[i].i = i
	}
	slices.SortFunc(a, func(a, b pair) int { return b.v - a.v })

	ans := make([][]byte, n+1)
	for i := range ans {
		ans[i] = bytes.Repeat([]byte{'0'}, n)
	}
	for i, p := range a {
		for j := 0; j < p.v; j++ {
			ans[(i+j)%(n+1)][p.i] = '1'
		}
	}
	Fprintln(out, n+1)
	for _, row := range ans {
		Fprintf(out, "%s\n", row)
	}

}

func main() { CF1227G(os.Stdin, os.Stdout) }
