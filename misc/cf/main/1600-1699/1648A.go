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

func CF1648A(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, m, v, ans int
	Fscan(in, &n, &m)
	r, c := make(map[int][]int), make(map[int][]int)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			Fscan(in, &v)
			r[v] = append(r[v], i)
			c[v] = append(c[v], j)
		}
	}

	f := func(val []int) int{
		s, res := 0, 0
		for i, v := range val {
			res += i * v - s
			s += v
		}
		return res
	}

	for _, v := range r {
		ans += f(v)
	}
	for _, v := range c {
		sort.Ints(v)
		ans += f(v)
	}
	Fprintln(out, ans)
}

func main() { CF1648A(os.Stdin, os.Stdout) }
