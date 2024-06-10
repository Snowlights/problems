//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"sort"
)

func CF991C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n int
	Fscan(in, &n)
	Fprint(out, 1+sort.Search(n/2, func(k int) bool {
		k++
		m := n
		s := 0
		for m > k {
			s += k
			m -= k
			m -= m / 10
		}
		return (s+m)*2 >= n
	}))
}

func main() { CF991C(os.Stdin, os.Stdout) }
