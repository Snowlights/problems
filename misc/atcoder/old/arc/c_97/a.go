package c_97

import (
	"bufio"
	"fmt"
	"io"
)

// https://atcoder.jp/contests/arc097/tasks/arc097_b
//
//输入 n(≤1e5) m(≤1e5) 和一个 1~n 的排列 p。
//然后输入 m 行，每行两个数 x 和 y，
//表示你可以交换 p[x] 和 p[y]（下标从 1 开始）。
//这 m 个操作你可以按任意顺序执行任意多次。
//输出你可以让多少个 p[i]=i。

func AtCoderARC97B(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	var n, m, x, y int
	fmt.Fscan(r, &n, &m)
	p := make([]int, n)
	for i := range p {
		fmt.Fscan(r, &p[i])
	}

	fa := make([]int, n+1)
	for i := range fa {
		fa[i] = i
	}
	var find func(int) int
	find = func(from int) int {
		if fa[from] != from {
			fa[from] = find(fa[from])
		}
		return fa[from]
	}
	merge := func(from, to int) {
		from, to = find(from), find(to)
		if from == to {
			return
		}
		fa[from] = to
	}

	for i := 0; i < m; i++ {
		fmt.Fscan(r, &x, &y)
		merge(x, y)
	}

	ans := 0
	for i, v := range p {
		if find(v) == find(i+1) {
			ans++
		}
	}
	fmt.Fprintln(w, ans)
}
