//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"slices"
)

func CF1012A(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n int
	Fscan(in, &n)
	a := make([]int, n*2)
	for i := range a {
		Fscan(in, &a[i])
	}
	slices.Sort(a)
	ans := (a[n-1] - a[0]) * (a[n*2-1] - a[n])
	for i := n; i < n*2-1; i++ {
		ans = min(ans, (a[i]-a[i-n+1])*(a[n*2-1]-a[0]))
	}
	Fprint(out, ans)

}

func main() { CF1012A(os.Stdin, os.Stdout) }
