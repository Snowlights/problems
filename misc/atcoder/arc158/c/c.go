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

	var n, mx, ds, carry int
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
		mx = max(mx, a[i])
		for x := a[i]; x > 0; x /= 10 {
			ds += x % 10
		}
	}
	for pow := 10; pow <= mx*2; pow *= 10 {
		sort.Slice(a, func(i, j int) bool {
			return a[i]%pow < a[j]%pow
		})
		j := n - 1
		for _, v := range a {
			for j >= 0 && v%pow+a[j]%pow >= pow {
				j--
			}
			carry += n - 1 - j
		}
	}

	Fprintln(out, 2*n*ds-9*carry)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() { run(os.Stdin, os.Stdout) }
