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
	a := make([][]int, n)
	for i := range a {
		a[i] = make([]int, 3)
		Fscan(in, &a[i][0], &a[i][1], &a[i][2])
		sort.Ints(a[i])
	}
	sort.Slice(a, func(i int, j int) bool {
		return a[i][0] < a[j][0]
	})

	f := &fenwick{}
	for i := 0; i < n; {
		st, val := i, a[i][0]
		for ; i < n && a[i][0] == val; i++ {
			if f.pre(a[i][1]-1) < a[i][2] {
				Fprintln(out, "Yes")
				return
			}
		}
		for ; st < i; st++ {
			f.upd(a[st][1], a[st][2])
		}
	}

	Fprintln(out, "No")
}

type fenwick map[int]int

func (f fenwick) pre(i int) int {
	res := int(1e9)
	for ; i > 0; i &= i - 1 {
		if v, ok := f[i]; ok && v < res {
			res = v
		}
	}
	return res
}

func (f fenwick) upd(i, val int) {
	for ; i <= 1e9; i += i & -i {
		if v, ok := f[i]; !ok || val < v {
			f[i] = val
		}
	}
}

func main() { run(os.Stdin, os.Stdout) }
