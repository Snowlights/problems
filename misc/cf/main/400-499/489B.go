//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"sort"
)

func CF489B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, m, ans, j int
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	Fscan(in, &m)
	b := make([]int, m)
	for i := range b {
		Fscan(in, &b[i])
	}

	sort.Ints(a)
	sort.Ints(b)
	for _, v := range a {
		for j < m && b[j] < v-1 {
			j++
		}
		if j < m && b[j] <= v+1 {
			ans++
			j++
		}
	}

	Fprint(out, ans)
}

func main() { CF489B(os.Stdin, os.Stdout) }
