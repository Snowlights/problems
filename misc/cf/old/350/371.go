package _50

import (
	"bufio"
	"fmt"
	"io"
)

func CF371D(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	var n, m int
	fmt.Fscan(r, &n)
	cp, water := make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(r, &cp[i])
	}
	var op, x, idx int

	fa := make([]int, n)
	for i := range fa {
		fa[i] = i
	}
	var find func(int) int
	find = func(i int) int {
		if fa[i] != i {
			fa[i] = find(fa[i])
		}
		return fa[i]
	}

	for fmt.Fscan(r, &m); m > 0; m-- {
		fmt.Fscan(r, &op)
		switch op {
		case 1:
			fmt.Fscan(r, &idx, &x)
			idx--
			for idx = find(idx); ; idx = find(idx) {
				if water[idx]+x < cp[idx] {
					water[idx] += x
					break
				}
				x -= cp[idx] - water[idx]
				water[idx] = cp[idx]
				idx++
				if idx == n {
					break
				}
				fa[find(idx-1)] = find(idx)
			}

		case 2:
			fmt.Fscan(r, &idx)
			idx--
			fmt.Fprintln(w, water[idx])
		}

	}

}
