package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, ans int
	const mod int = 1e9 + 7
	Fscan(in, &n)
	a, b := make([]int, n), make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	for i := range b {
		Fscan(in, &b[i])
	}

	ans = 1
	for i, av := range a {
		af := i == 0 || a[i] > a[i-1]
		bf, bv := i == n-1 || b[i] > b[i+1], b[i]

		if af && av > bv || bf && bv > av {
			Fprintln(out, 0)
			return
		}
		if !af && !bf {
			ans = ans * min(av, bv) % mod
		}
	}
	Fprintln(out, ans)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() { run(os.Stdin, os.Stdout) }
