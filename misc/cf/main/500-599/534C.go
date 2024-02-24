//go:build main
// +build main

package main



import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF534C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, s, sa int
	Fscan(in, &n, &s)
	s -= n
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
		a[i]--
		sa += a[i]
	}
	for _, v := range a {
		mn := max(s-sa+v, 0)
		mx := min(s, v)
		Fprint(out, v-mx+mn, " ")
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}


func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}


func main() { CF534C(os.Stdin, os.Stdout) }
