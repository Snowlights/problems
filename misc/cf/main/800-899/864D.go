//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF864D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, ans int
	Fscan(in, &n)
	a, h, cur, vis := make([]int, n), make(map[int]int), 1, make([]bool, n)
	for i := range a {
		Fscan(in, &a[i])
		h[a[i]]++
	}
	for i, v := range a {
		if h[v] == 1 {
			continue
		}
		for h[cur] > 0 {
			cur++
		}
		if cur > v && !vis[v] {
			vis[v] = true
			continue
		}
		h[v]--
		h[cur]++
		a[i] = cur
		ans++
	}
	Fprintln(out, ans)
	for _, v := range a {
		Fprint(out, v, " ")
	}
}

func main() { CF864D(os.Stdin, os.Stdout) }
