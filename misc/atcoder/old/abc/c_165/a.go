package c_165

import (
	"bufio"
	"fmt"
	"io"
	"sort"
)

// https://atcoder.jp/contests/abc165/tasks/abc165_f
//
// 输入 n (2≤n≤2e5) 和长为 n 的数组 a (1≤a[i]≤1e9)，表示每个节点的点权。
// 然后输入一棵树的 n-1 条边（节点编号从 1 开始）。
// 输出 n 个数，第 i 个数为从节点 1 到节点 i 的路径上点权的 LIS 长度。
//
//注：LIS 指最长上升子序列。

func AtCoderABC165F(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	var n, u, v int
	fmt.Fscan(r, &n)
	a := make([]int, n)
	for i := range a {
		fmt.Fscan(r, &a[i])
	}
	g := make([][]int, n)
	for i := 0; i < n-1; i++ {
		fmt.Fscan(r, &u, &v)
		v--
		u--
		g[u] = append(g[u], v)
		g[v] = append(g[v], u)
	}

	lis, ans := make([]int, 0), make([]int, n)
	var dfs func(fa, v int)
	dfs = func(fa, v int) {
		if i := sort.SearchInts(lis, a[v]); i < len(lis) {
			old := lis[i]
			lis[i] = a[v]
			ans[v] = len(lis)
			for _, to := range g[v] {
				if to == fa {
					continue
				}
				dfs(v, to)
			}
			lis[i] = old
		} else {
			lis = append(lis, a[v])
			ans[v] = len(lis)
			for _, to := range g[v] {
				if to == fa {
					continue
				}
				dfs(v, to)
			}
			lis = lis[:len(lis)-1]
		}
	}
	dfs(-1, 0)
	for _, v := range ans {
		fmt.Fprintln(w, v)
	}

}
