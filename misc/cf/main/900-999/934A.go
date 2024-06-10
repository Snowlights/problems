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

func CF934A(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, m int
	Fscan(in, &n, &m)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	sort.Ints(a)
	b := make([]int, m)
	for i := range b {
		Fscan(in, &b[i])
	}
	sort.Ints(b)

	if max(a[0]*b[0], a[0]*b[m-1]) > max(a[n-1]*b[0], a[n-1]*b[m-1]) {
		a = a[1:]
	} else {
		a = a[:n-1]
	}
	n--
	Fprint(out, max(max(a[0]*b[0], a[0]*b[m-1]), max(a[n-1]*b[0], a[n-1]*b[m-1])))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() { CF934A(os.Stdin, os.Stdout) }
