package _50

import (
	"bufio"
	"fmt"
	"io"
)

func CF274B(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	var n, u, v int
	fmt.Fscan(r, &n)
	g := make(map[int][]int)
	for i := 0; i < n-1; i++ {
		fmt.Fscan(r, &u, &v)
		u--
		v--
		g[u] = append(g[u], v)
		g[v] = append(g[v], u)
	}
	arr := make([]int, n)
	for i := range arr {
		fmt.Fscan(r, &arr[i])
	}

	min := func(a, b int64) int64 {
		if a > b {
			return b
		}
		return a
	}
	max := func(a, b int64) int64 {
		if a < b {
			return b
		}
		return a
	}

	var f func(int, int) (int64, int64)
	f = func(v, fa int) (mn, mx int64) {
		cur := int64(arr[v])
		if len(g[v]) == 0 { // 叶子
			return min(cur, 0), max(cur, 0)
		}
		for _, w := range g[v] {
			if w != fa {
				x, y := f(w, v)
				mn = min(mn, x)
				mx = max(mx, y)
			}
		}
		cur -= mn + mx
		if cur < 0 {
			mn += cur
		} else {
			mx += cur
		}
		return
	}

	mn, mx := f(0, -1)
	fmt.Fprintln(w, mx-mn)
}
