package _450

import (
	"bufio"
	"fmt"
	"io"
	"sort"
)

func CF1494C(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	f := func(a, b []int) (mx int) {
		h := map[int]bool{}
		for _, v := range a {
			h[v] = true
		}
		s := 0
		for _, v := range b {
			if h[v] {
				s++
			}
		}
		for i, v := range b {
			if h[v] {
				s--
			}
			j := sort.SearchInts(b, v-sort.SearchInts(a, v+1)+1)
			if i-j+1+s > mx {
				mx = i - j + 1 + s
			}
		}
		return
	}
	var T, n, m int64
	for fmt.Fscan(r, &T); T > 0; T-- {
		fmt.Fscan(r, &n, &m)
		a, b := make([]int, n), make([]int, m)
		for i := range a {
			fmt.Fscan(r, &a[i])
		}
		for i := range b {
			fmt.Fscan(r, &b[i])
		}
		x, y := sort.SearchInts(a, 0), sort.SearchInts(b, 0)
		c := f(a[x:], b[y:])
		for i := 0; i < (x+1)/2; i++ {
			a[i], a[x-1-i] = -a[x-1-i], -a[i]
		}
		for i := 0; i < (y+1)/2; i++ {
			b[i], b[y-1-i] = -b[y-1-i], -b[i]
		}
		c += f(a[:x], b[:y])
		fmt.Fprintln(out, c)
	}

}
