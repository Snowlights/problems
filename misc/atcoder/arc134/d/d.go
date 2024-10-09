package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"sort"
)

func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n int
	Fscan(in, &n)
	a, b := make([]int, n), make([]int, n)
	x, y := make([]int, 0, n), make([]int, 0, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	for i := range b {
		Fscan(in, &b[i])
	}

	for i, v := range a {
		for len(x) > 0 && v < x[len(x)-1] {
			x, y = x[:len(x)-1], y[:len(y)-1]
		}
		x, y = append(x, v), append(y, b[i])
	}

	mn := y[0]
	for i, v := range x {
		if v > x[0] {
			break
		}
		mn = min(mn, y[i])
	}
	if mn <= x[0] {
		Fprint(out, x[0], mn)
		return
	}

	r := sort.SearchInts(x, y[0]+1)
	x = x[:r]
	l := sort.SearchInts(x, y[0])

	ans := append(x, y[:r]...)
	less := func() bool {
		for i, v := range y[:l] {
			if v != ans[l+i] {
				return v < ans[l+i]
			}
		}
		return true
	}
	if less() {
		ans = append(ans[:l], ans[r:r+l]...)
	}

	for _, v := range ans {
		Fprint(out, v, " ")
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() { run(os.Stdin, os.Stdout) }
