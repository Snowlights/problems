package _350

import (
	"bufio"
	"fmt"
	"io"
)

func CF1395C(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	var n, m int
	fmt.Fscan(r, &n, &m)
	a, b := make([]int, n), make([]int, m)
	for i := range a {
		fmt.Fscan(r, &a[i])
	}
	for i := range b {
		fmt.Fscan(r, &b[i])
	}

O:
	for ans := 0; ; ans++ {
	o:
		for _, av := range a {
			for _, bv := range b {
				if av&bv|ans == ans {
					continue o
				}
			}
			continue O
		}
		fmt.Fprintln(w, ans)
		return
	}
}
