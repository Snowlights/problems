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
	a := make([]float64, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	sort.Slice(a, func(i, j int) bool {
		return 1/a[i] < 1/a[j]
	})
	f := func(a, b, c float64) float64 {
		return (a + b + c) / (a * b * c)
	}
	v1, v2, v3, v4 := f(a[n-1], a[n-2], a[n-3]), f(a[n-1], a[n-2], a[0]), f(a[n-1], a[0], a[1]), f(a[0], a[1], a[2])
	Fprintln(out, Sprintf("%.15f", min(v1, v2, v3, v4)))
	Fprintln(out, Sprintf("%.15f", max(v1, v2, v3, v4)))

}

func max(a ...float64) float64 {
	res := a[0]
	for _, v := range a {
		if v > res {
			res = v
		}
	}
	return res
}

func min(a ...float64) float64 {
	res := a[0]
	for _, v := range a {
		if v < res {
			res = v
		}
	}
	return res
}

func main() { run(os.Stdin, os.Stdout) }
