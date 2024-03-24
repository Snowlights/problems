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

func CF817B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, v int
	Fscan(in, &n)
	cnt := make(map[int]int)
	for i := 0; i < n; i++ {
		Fscan(in, &v)
		cnt[v]++
	}
	a := make([]int, 0, len(cnt))
	for k := range cnt {
		a = append(a, k)
	}
	sort.Ints(a)

	c0 := cnt[a[0]]
	if c0 > 2 {
		Fprintln(out, c0 * (c0-1) * (c0-2) / 6)
	} else if c1 := cnt[a[1]]; c0 == 2 {
		Fprintln(out, c0 * (c0-1) / 2 * c1)
	} else if c1 > 1 {
		Fprintln(out,  c1 * (c1 - 1) / 2 * c0)
	} else {
		Fprintln(out, c0 * c1 * cnt[a[2]])
	}
	
}

func main() { CF817B(os.Stdin, os.Stdout) }
