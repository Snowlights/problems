package _0

import (
	"bufio"
	"fmt"
	"io"
	"sort"
)

// https://codeforces.com/problemset/problem/38/E
//
//输入 n(≤3000)，表示在一维数轴上有 n 颗大小忽略不计的弹珠。
//然后输入 n 对数字，每对数字表示这颗弹珠在数轴上的位置 xi，
// 以及把这颗弹珠固定在 xi 上的花费 ci，数据范围均在 [-1e9,1e9] 之间，
// 且 xi 互不相同。注意 ci 可以为负。
//
//选择若干弹珠固定后，所有未固定的弹珠向左滚动，直到碰到固定住的弹珠。
//总花费 = 固定弹珠的花费之和 + 所有未固定的弹珠的滚动距离之和。
//输出总花费的最小值。
//注：根据题意，最左边的弹珠一定要固定。

func CF38E(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	var n int
	fmt.Fscan(r, &n)
	type pair struct {
		x, c int64
	}
	pList := make([]pair, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(r, &pList[i].x, &pList[i].c)
	}
	sort.Slice(pList, func(i, j int) bool {
		return pList[i].x < pList[j].x
	})
	min := func(a, b int64) int64 {
		if a > b {
			return b
		}
		return a
	}
	//定义 f[i][j] 表示考虑前 i 颗弹珠，且最后一个固定的弹珠是 j 时的最小花费，则有：
	// f[i][j] = f[i-1][j] + x[i] - x[j],  j < i
	// f[i][i] = min(f[i-1][j]) + c[i]
	// 初始 f[0][0] = c[0]，
	// 答案 min(f[n-1])。
	dp := make([][]int64, n)
	for i := range dp {
		dp[i] = make([]int64, n)
	}
	dp[0][0] = pList[0].c
	for i := 1; i < n; i++ {
		mm := dp[i-1][0]
		for j := 0; j < i; j++ {
			dp[i][j] = dp[i-1][j] + pList[i].x - pList[j].x
			mm = min(mm, dp[i-1][j])
		}
		dp[i][i] = mm + pList[i].c
	}
	ans := dp[n-1][0]
	for _, v := range dp[n-1] {
		ans = min(ans, v)
	}
	fmt.Fprintln(w, ans)
}
