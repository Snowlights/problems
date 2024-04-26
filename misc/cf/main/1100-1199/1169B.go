//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF1169B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n int
	Fscan(in, &n, &n)
	a := make([]struct{ x, y int }, n)
	for i := range a {
		Fscan(in, &a[i].x, &a[i].y)
	}

	var f func(int, int) bool
	f = func(x, y int) bool {
		for _, p := range a {
			if p.x == x || p.x == y || p.y == x || p.y == y {
				continue
			}
			if y > 0 {
				return false
			}
			return f(x, p.x) || f(x, p.y)
		}
		return true
	}
	if f(a[0].x, 0) || f(a[0].y, 0) {
		Fprintln(out, "YES")
	} else {
		Fprintln(out, "NO")
	}

}

func main() { CF1169B(os.Stdin, os.Stdout) }
