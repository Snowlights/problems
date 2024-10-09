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

	var n, m, ans int
	Fscan(in, &n, &n, &m)
	type pair struct {
		v, l int
	}
	a, b := make([]pair, n), make([]pair, m)
	for i := range a {
		Fscan(in, &a[i].v, &a[i].l)
	}
	for i := range b {
		Fscan(in, &b[i].v, &b[i].l)
	}

	for i, j := 0, 0; i < n && j < m; {
		if a[i].v == b[j].v {
			ans += min(a[i].l, b[j].l)
		}
		if a[i].l > b[j].l {
			a[i].l -= b[j].l
			j++
		} else {
			b[j].l -= a[i].l
			i++
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
