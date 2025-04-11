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

	var n, q, l, r int
	Fscan(in, &n, &q)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	type pair struct{ r, i int }
	ls := make([][]pair, n)
	for i := 0; i < q; i++ {
		Fscan(in, &l, &r)
		ls[l] = append(ls[l], pair{r - 1, i})
	}

	ans := make([]int, q)
	st := []int{}
	for i := n - 1; i > 0; i-- {
		v := a[i]
		for len(st) > 0 && v > a[-st[len(st)-1]] {
			st = st[:len(st)-1]
		}
		st = append(st, -i)
		for _, p := range ls[i] {
			ans[p.i] = sort.SearchInts(st, -p.r)
		}
	}
	for _, v := range ans {
		Fprintln(out, v)
	}

}

func main() { run(os.Stdin, os.Stdout) }
