package c_266

import (
	"bufio"
	"fmt"
	"io"
)

// https://atcoder.jp/contests/abc266/tasks/abc266_f
//
//输入 n (3≤n≤2e5) 和 n 条边，点的编号在 [1,n] 内，表示一个没有重边和自环的无向连通图。
//然后输入 q(≤2e5) 表示有 q 个询问，每个询问输入两个数 x 和 y (1≤x<y≤n)。
//对于每个询问，如果 x 和 y 之间只存在唯一的简单路径，则输出 Yes，否则输出 No。

func AtCoderABC266F(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	var n, m, v, u, id int
	fmt.Fscan(in, &n)
	g := make([][]int, n)
	deg := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &v, &u)
		v--
		u--
		g[v] = append(g[v], u)
		g[u] = append(g[u], v)
		deg[v]++
		deg[u]++
	}

	q := []int{}
	for i, d := range deg {
		if d == 1 {
			q = append(q, i)
		}
	}
	for len(q) > 0 {
		v := q[0]
		q = q[1:]
		for _, w := range g[v] {
			if deg[w]--; deg[w] == 1 {
				q = append(q, w)
			}
		}
	}

	ids := make([]int, n)
	var dfs func(int, int)
	dfs = func(v, fa int) {
		ids[v] = id
		for _, w := range g[v] {
			if w != fa && deg[w] < 2 {
				dfs(w, v)
			}
		}
	}
	for i, d := range deg {
		if d > 1 {
			id++
			dfs(i, -1)
		}
	}

	for fmt.Fscan(r, &m); m > 0; m-- {
		fmt.Fscan(r, &v, &u)
		if ids[v-1] == ids[u-1] {
			fmt.Fprintln(w, "Yes")
		} else {
			fmt.Fprintln(w, "No")
		}
	}

}
