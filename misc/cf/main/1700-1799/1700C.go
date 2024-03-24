//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF1700C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		s, ans := a[0], 0
		for i := 1; i < n; i++ {
			d := a[i] - a[i-1]
			ans += abs(d)
			if d < 0 {
				s += d
			}
		}
		ans += abs(s)
		Fprintln(out, ans)
	}

}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() { CF1700C(os.Stdin, os.Stdout) }
