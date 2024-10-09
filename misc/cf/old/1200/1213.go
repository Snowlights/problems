package _200

import (
	"bufio"
	"fmt"
	"io"
	"sort"
)

// https://codeforces.com/problemset/problem/1213/G
//
//输入 n(≤2e5) 和 m(≤2e5)，表示一棵有 n 个节点的树，和 m 个询问。
//然后输入 n-1 条边：这条边所连接的两点的编号（从 1 开始），以及边权(1≤边权≤2e5)。
//最后输入 m 个询问 q[i](1≤q[i]≤2e5)，你需要对每个询问，输出树上有多少条路径，
//要求路径至少有两个节点且无重复节点，且路径上的最大边权不超过 q[i]。
//注意 a->b 和 b->a 的路径只统计一次。
//
//相关题目：昨天周赛第四题 https://leetcode.cn/problems/number-of-good-paths/

func CF1213G(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	var n, m int
	fmt.Fscan(r, &n, &m)
	type node struct {
		from, to, weight int
	}
	type query struct {
		q, i int
	}
	nodeList := make([]node, n-1)
	queryList := make([]query, m)
	for i := range nodeList {
		fmt.Fscan(r, &nodeList[i].from, &nodeList[i].to, &nodeList[i].weight)
	}
	for i := range queryList {
		fmt.Fscan(r, &queryList[i].q)
		queryList[i].i = i
	}
	sort.Slice(nodeList, func(i, j int) bool {
		return nodeList[i].weight < nodeList[j].weight
	})
	sort.Slice(queryList, func(i, j int) bool {
		return queryList[i].q < queryList[j].q
	})

	// 并查集模版
	fa := make([]int, n+1)
	size := make([]int, n+1)
	for i := range fa {
		fa[i] = i
		size[i] = 1
	}
	var find func(int) int
	find = func(x int) int {
		if fa[x] != x {
			fa[x] = find(fa[x])
		}
		return fa[x]
	}
	ans := make([]int64, m)
	c, i := int64(0), 0
	for _, q := range queryList {
		for ; i < n-1 && nodeList[i].weight <= q.q; i++ {
			from, to := find(nodeList[i].from), find(nodeList[i].to)
			c += int64(size[from] * size[to])
			size[to] += size[from]
			fa[from] = to
		}
		ans[q.i] = c
	}
	for _, v := range ans {
		fmt.Fprint(w, v, " ")
	}

}
