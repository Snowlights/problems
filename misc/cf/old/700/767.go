package _00

import (
	"bufio"
	"fmt"
	"io"
)

// https://codeforces.com/problemset/problem/767/C
//
//输入 n(3≤n≤1e6)，表示一颗有 n 个节点的有根树，节点编号从 1 开始。
//每行输入两个数，表示当前节点的父节点编号（如果是 0 表示当前节点是根节点），以及节点的点权，范围在 [-100,100]。
//例如节点 1 的父节点为 2，则表示一条 2->1 的边。
//
//你需要删除两条边，将这棵树分成三个连通块，且每个连通块的点权和相等。
//假设删除的边是 a->b 和 c->d，你需要输出 b 和 d。如果有多种方案，输出任意一种。
//如果无法做到，输出 -1。

var f [1e6 + 5]int64
var g [1e6 + 5][]int64
var total, root int64
var x, y int64

func dfs(son int64) int64 {
	val := f[son]
	for _, v := range g[son] {
		val += dfs(v)
	}

	if son != root && val == total {
		if x == 0 {
			x = son
		} else if y == 0 {
			y = son
		}
		val = 0
	}
	return val
}

func CF767C(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	var n, to int64
	fmt.Fscan(r, &n)
	for i := int64(1); i <= n; i++ {
		fmt.Fscan(r, &to, &f[i])
		total += f[i]
		if to == 0 {
			root = i
		} else {
			g[to] = append(g[to], i)
		}
	}
	if total%3 != 0 {
		fmt.Fprintln(w, -1)
		return
	}
	total = total / 3
	dfs(root)
	// fmt.Println(pairList)
	if y != 0 {
		fmt.Fprintln(w, x, y)
	} else {
		fmt.Fprintln(w, -1)
	}

}
