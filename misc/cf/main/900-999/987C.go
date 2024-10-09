//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"math"
	"os"
)

func CF987C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n int
	Fscan(in, &n)
	a, b := make([]int, n), make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	for i := range b {
		Fscan(in, &b[i])
	}
	ans := math.MaxInt32
	for i := 1; i < n-1; i++ {
		l := math.MaxInt32
		for j := 0; j < i; j++ {
			if a[j] < a[i] {
				l = min(l, b[j])
			}
		}
		r := math.MaxInt32
		for k := i + 1; k < n; k++ {
			if a[k] > a[i] {
				r = min(r, b[k])
			}
		}
		ans = min(ans, l+b[i]+r)
	}
	if ans == math.MaxInt32 {
		Fprintln(out, -1)
		return
	}
	Fprintln(out, ans)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() { CF987C(os.Stdin, os.Stdout) }
