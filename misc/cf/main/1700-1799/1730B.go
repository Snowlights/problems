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

func CF1730B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, t int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		mm, mx := math.MaxInt32, 0
		for i := range a {
			Fscan(in, &t)
			mm, mx = min(mm, a[i]-t), max(mx, a[i]+t)
		}
		Fprintln(out, float64(mm+mx)/2)
	}

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() { CF1730B(os.Stdin, os.Stdout) }
